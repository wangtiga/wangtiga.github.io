---
layout: post
title:  "gRPC 服务性能验证"
date:   2023-04-15 12:00:00 +0800
tags:   tech
---

* category
{:toc}


## QPS 预估

grpc 服务在 8C8G 主机上的 qps，约 15wQPS [^grpcGrafanaDotGrpcTesting], 其中 2022.01 的 async API C++ 版本 v1.4x 版本性能最好,约 20wQPS;


预估最大 QPS 的方法[^maxQPSPlantegg]

- 单线程QPS = 1000ms / RT(Thread CPU Time + Thread Wait Time)
- 最佳线程数 = (RT/CPU Time) x CPU核数 x CPU利用率
- 最大QPS = (1000ms/CPU Time) x CPU核数 x CPU利用率


以 RT=100ms, CPU核数=8Core, CPU利用率=50% 为例,
- 单线程QPS = 1000ms/100ms = 10;
- 最佳线程数 = (10/1ms) x 8 x 0.5 = 40;
- 最大QPS = (1000ms/1ms) x 8 x 0.5 = 4000;

> CPU利用率与代码实现有关,有全局大锁会导致利用率低,具体取值要频繁观测得出；

> 如果已知当前最大 2wQPS , 50%的8核心CPU,则反推实际 CPU Time 是 (1000ms/20000) x 8 x 0.5 = (1/20) x 4 = 0.05 x 4 = 0.2ms = 200us



## async API 使用示例

### 0. 参考 grpc 官方 example/cpp/helloworld 熟悉 async API

https://github.com/wangtiga/grpc/blob/v1.44-dev/examples/cpp/helloworld/greeter_async_server.cc

### 1. 在 async server 增加多线程  gretter_async_server_mthread 

- add multithread 多线程处理队列请求
```c
  void ServerImpl::HandleRpcsMultiThread() {
    std::vector<std::future<void>> futures;
    int num_threads = std::thread::hardware_concurrency();
    std::cout << "hardware_concurrency thread num " << num_threads << std::endl;
    for (int i = 0; i < num_threads; ++i) {
      futures.push_back(std::async(std::launch::async, [this]() {
        HandleRpcs();
      }));
    }   

    for (auto&& f : futures) {
      f.wait();
    }   
  }


  // This can be run in multiple threads if needed.
  void ServerImpl::HandleRpcs() {
    new CallData(&service_, cq_.get());
    void* tag;  // uniquely identifies a request.
    bool ok; 
    while (true) {
      // Block waiting to read the next event from the completion queue. The
      // event is uniquely identified by its tag, which in this case is the
      // memory address of a CallData instance.
      // The return value of Next should always be checked. This return value
      // tells us whether there is any kind of event or cq_ is shutting down.
      GPR_ASSERT(cq_->Next(&tag, &ok));
      GPR_ASSERT(ok);
      static_cast<CallData*>(tag)->Proceed();
    }
  }
```

### 2. 在 async server 中增加 async call 调用, 模拟外部请求 gretter_async_server_mthread2

https://github.com/wangtiga/grpc/blob/v1.44-dev/examples/cpp/helloworld/greeter_async_server_mthread2.cc

- external async call 增加异步的外部接口调用

```c
    void CallData::Proceed() {
      if (status_ == CREATE) {
        status_ = PROCESS1;
        service_->RequestSayHello(&ctx_, &request_, &responder_, cq_, cq_,
                                  this);
      } else if (status_ == PROCESS1) {
        status_ = PROCESS2;
        client_rpc_->SayHello("rpc1", this);
      } else if (status_ == PROCESS2) {
        status_ = PROCESS_LAST;
        client_rpc_->SayHello("rpc2", this);
      } else if (status_ == PROCESS_LAST) {
        new CallData(service_, cq_, client_rpc_);

        status_ = FINISH;

        prefix = prefix + " " + request_.name();
        prefix = prefix + " sub_reply1: " + sub_reply1_.message();
        prefix = prefix + " sub_reply2: " + sub_reply2_.message();
        reply_.set_message(prefix);
        responder_.Finish(reply_, Status::OK, this);
      } else {
        GPR_ASSERT(status_ == FINISH);
        delete this;
      }
    }
```


### 3. 将 client server 共用一个 CompletionQueue 并预留足够的请求数 gretter_async_server_mthread3

https://github.com/wangtiga/grpc/blob/v1.44-dev/examples/cpp/helloworld/greeter_async_server_mthread2.cc


- server 和 client 共用一个 CompletionQueue 队列,以减少循环检查队列的线程

// helloworld.grpc.pb.h
```c
// server 的 `Greeter::AsyncService.RequestSayHello` 方法；
void RequestSayHello(::grpc::ServerContext* context, ::helloworld::HelloRequest* request, ::grpc::ServerAsyncResponseWriter< ::helloworld::HelloReply>* response, ::gr    pc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag

// client 的 `Greeter::Stub.PrepareAsyncSayHello` 方法;
std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::helloworld::HelloReply>> PrepareAsyncSayHello(::grpc::ClientContext* context, const ::helloworld::HelloRequest&     request, ::grpc::CompletionQueue* cq) {
```

```c
    void CallData::ProceedSendGreeterSayHello(const std::string& user) {
      HelloRequest request;
      request.set_name(user);

      // stub_->PrepareAsyncSayHello() creates an RPC object, returning
      // an instance to store in "call" but does not actually start the RPC
      // Because we are using the asynchronous API, we need to hold on to
      // the "call" instance in order to get updates on the ongoing RPC.
      sub_response_reader1_ =
          client_rpc_->stub_->PrepareAsyncSayHello(&client_rpc_context1_, request, cq_);

      // StartCall initiates the RPC call
      sub_response_reader1_->StartCall();

      // Request that, upon completion of the RPC, "reply" be updated with the
      // server's response; "status" with the indication of whether the operation
      // was successful. Tag the request with the memory address of the call
      // object.
      sub_response_reader1_->Finish(&sub_reply1_, &sub_status1_, (void*)this);
    }
```

```c
    void CallData::Proceed() {
      if (status_ == CREATE) {
        status_ = PROCESS1;

        // As part of the initial CREATE state, we *request* that the system
        // start processing SayHello requests. In this request, "this" acts are
        // the tag uniquely identifying the request (so that different CallData
        // instances can serve different requests concurrently), in this case
        // the memory address of this CallData instance.
        service_->RequestSayHello(&ctx_, &request_, &responder_, cq_, cq_,
                                  this);
      } else if (status_ == PROCESS1) {
    // other code ...
```

- CompletionQueue 中预留请求数,以增大并发

```c
  void ServerImpl::HandleRpcs(ServerCompletionQueue* cq, GreeterClient* client_rpc, int max_calldata_in_queue) {
    std::cout << "HandleRpcs max_calldata_in_queue: " << max_calldata_in_queue << std::endl;
    for (int i = 0; i < max_calldata_in_queue; ++i) {
      // Spawn a new CallData instance to serve new clients.
      new CallData(&service_, cq, client_rpc);
    }   
    void* tag;
    bool ok; 
    while (true) {
      GPR_ASSERT(cq->Next(&tag, &ok));
      GPR_ASSERT(ok);
      static_cast<CallData*>(tag)->Proceed();
    }   
  }
```

```c
  // Class encompasing the state and logic needed to serve a request.
  class CallData : public CallProceeder {
   public:
    CallData(Greeter::AsyncService* service, ServerCompletionQueue* cq, GreeterClient* client_rpc)
        : service_(service), cq_(cq), responder_(&ctx_), status_(CREATE), client_rpc_(client_rpc) {
      Proceed();
    }
    // other code ...
```



其他参数 ：
1. MAX_CONCURRENT_STREAM  
2. SyncServerOption::MAX_POLLERS
3. GRPC_ARG_MAX_CONCURRENT_STREAMS


## C++ gRPC 性能最佳实践 [^gRPCPerformanceBestPractices]

> Do not use Sync API for performance sensitive servers. If performance and/or resource consumption are not concerns, use the Sync API as it is the simplest to implement for low-QPS services.

同步API使用简单，只适用于低QPS的服务。
如果对性能有要求，不要用 Sync API；

> Favor callback API over other APIs for most RPCs, given that the application can avoid all blocking operations or blocking operations can be moved to a separate thread. The callback API is easier to use than the completion-queue async API but is currently slower for truly high-QPS workloads.

- 使用不同类型 API 能实现的 QPS 从高到低排列: 1 2 3 ;
- 易用程度从高到低排序: 3 2 1;

1. completion-queue async API
2. callback API 
3. other APIs （ Sync API）


> If having to use the async completion-queue API, the best scalability trade-off is having numcpu’s threads. The ideal number of completion queues in relation to the number of threads can change over time (as gRPC C++ evolves), but as of gRPC 1.41 (Sept 2021), using 2 threads per completion queue seems to give the best performance.

使用 async completion-queue API 时，线程数量与CPU核心数量相同，性能最好；
最佳队列数量与线程是不固定的，在 gRPC1.41 之后，2个线程共用1个completion-queue，性能最好；


> For the async completion-queue API, make sure to register enough server requests for the desired level of concurrency to avoid the server continuously getting stuck in a slow path that results in essentially serial request processing.

使用 async completion-queue API 时，在 completion-queue 中预备的 server request 数量要与实际并发请求量相当，避免运行效率恶化到与串行处理一样慢。

> 注：以官方demo greeter_async_server 为例，在启动服务时，多执行几次 service_->RequestSayHello() ，保证队列中有足够的 CallData ，满足实际的并发需要。默认队列中仅一个 CallData 时，实际的并发请求就是 1TPS ，即上一个 CallData 处理完毕，才有机会运行下一个 CallData；


> (Special topic) Enable write batching in streams if message k + 1 does not rely on responses from message k by passing a WriteOptions argument to Write with buffer_hint set: `stream_writer->Write(message, WriteOptions().set_buffer_hint());`

（专题）Stream RPC 中，如果第 k+1 个消息不依赖第 k 个消息的响应，可调用  set_buffer_hint 开启写缓冲来提高效率（write调用时不会立刻写入网卡发出去，而是根据数据量和网络状况决定当前消息是缓存还是立即发出）；

```c
// include/grpcpp/impl/codegen/call_op_set.h
/// Sets flag indicating that the write may be buffered and need not go out on                                                                                              
/// the wire immediately.                                                                                                                                              
inline WriteOptions& set_buffer_hint()
```


> (Special topic) gRPC::GenericStub can be useful in certain cases when there is high contention / CPU time spent on proto serialization. This class allows the application to directly send raw gRPC::ByteBuffer as data rather than serializing from some proto. This can also be helpful if the same data is being sent multiple times, with one explicit proto-to-ByteBuffer serialization followed by multiple ByteBuffer sends.

（专题）当在 proto 序列化操作上（编码、解码）消耗大量 CPU 时间或存在高并发请求时，gRPC::GenericStub 类可以派上用场。该类允许应用程序直接将原始的 gRPC::ByteBuffer 作为数据发送，而不必将某个 proto Message （如 helloworld::HelloRequest）序列化后再发出。如果需要多次发送相同的数据，则可以进行一次显式的 proto Message 到 ByteBuffer 序列化，然后重复发送 ByteBuffer 数据，这样也能提高效率。



[^gRPCPerformanceBestPractices]:[gRPC Performance Best Practices](https://grpc.io/docs/guides/performance/)

[^grpcGrafanaDotGrpcTesting]: [gRPC官方性能指标记录](https://grafana-dot-grpc-testing.appspot.com/?viewPanel=26&from=now-2y&to=now)

[^maxQPSPlantegg]: [性能优化，从老中医到科学理论指导-plantegg](https://plantegg.github.io/2018/08/24/%E6%80%A7%E8%83%BD%E4%BC%98%E5%8C%96%EF%BC%8C%E4%BB%8E%E8%80%81%E4%B8%AD%E5%8C%BB%E5%88%B0%E7%A7%91%E5%AD%A6%E7%90%86%E8%AE%BA%E6%8C%87%E5%AF%BC/)

