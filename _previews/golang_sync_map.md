

## sync.Map 

https://golang.org/pkg/sync/#Map

NOTE 不是所有情况都使用用 sync.Map 

Map is like a Go map[interface{}]interface{} but is safe for concurrent use by multiple goroutines without additional locking or coordination. Loads, stores, and deletes run in amortized constant time.

The Map type is specialized. Most code should use a plain Go map instead, with separate locking or coordination, for better type safety and to make it easier to maintain other invariants along with the map content.

The Map type is optimized for two common use cases: 

- (1) when the entry for a given key is only ever written once but read many times, as in caches that only grow, or
- (2) when multiple goroutines read, write, and overwrite entries for disjoint sets of keys. 

以下情况适合使用 sync.Map 

1. 写少读多的情况 
2. 多个 goroutine 读写,但互相之间 key 没有交集

In these two cases, use of a Map may significantly reduce lock contention compared to a Go map paired with a separate Mutex or RWMutex.

The zero Map is empty and ready for use. A Map must not be copied after first use.

```go
type Map struct {
    // contains filtered or unexported fields
}
```

## sync.Mutex

https://golang.org/pkg/sync/#Mutex

NOTE 值类型,发生拷贝后,会产生两个不同的 mutex 


A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.

A Mutex must not be copied after first use.


https://golang.org/src/sync/mutex.go?s=765:813#L15

```go
type Mutex struct {
	state int32
	sema  uint32
}
```

## TODO

https://blog.golang.org/race-detector

how to analyze panic output

