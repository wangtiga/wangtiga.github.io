## 1. 下面代码能运行吗？为什么

```go
type Param map[string]interface{}

type Show struct {
	Param
}

func main() {
	s := new(Show)
	s.Param["RMB"] = 10000
}
```

运行结果：
```shell
panic: assignment to entry in nil map

goroutine 1 [running]:
main.main()
```

如上所示，运行过程中会发生异常，
原因是因为字典Param的默认值为nil，当给字典nil增加键值对是就会发生运行时错误`panic: assignment to entry in nil map`。
另外，注意 `s := make(Show)` 也会有问题，因为 Param 是 Show 一个成员变量 。

> [https://wangtiga.github.io/golang/2019/05/31/effective-golang.html](https://wangtiga.github.io/golang/2019/05/31/effective-golang.html#allocation-with-make-%E4%BD%BF%E7%94%A8-make-%E5%88%86%E9%85%8D%E5%86%85%E5%AD%98)
> 内置函数`make(T, args)`与`new(T)`使用目的完全不同。它仅用于创建 slices, maps, channels ，并返回值是初始化过的（不是zero），且类型为`T`的变量（不是`*T`）。出现这种差异的本质(under the cover)原因是这三种类型是引用类型，必须在使用前初始化。比如，slice由三个descriptor（描述符）组成，分别指向 data （数组的数据），length （长度），capacity（容量），在三个descriptor未初始化前， slice 的值是 nil 。对于 slices, maps, channels 来说，`make`用于初始化结构体内部数据并赋值

正确的修改方案如下：
```go
package main

import "fmt"

type Param map[string]interface{}

type Show struct {
	Param
}

func main() {
	// 创建Show结构体对象
	s := new(Show)
	// 为字典Param赋初始值
	s.Param = Param{}
	// 修改键值对
	s.Param["RMB"] = 10000
	fmt.Println(s)
}
```

运行结果如下：

```shell
&{map[RMB:10000]}

Process finished with exit code 0
```

## 2. ~~请说出下面代码存在什么问题~~ 这题目不好，可以去掉了
```go
package main

import "fmt"

func main() {
    hello(student{"tiga"})
}

type student struct {
    Name string
}

func hello(v interface{}) {
    switch msg := v.(type) {
    case *student, student:
        fmt.Println(msg.Name)
    }
}
```

运行结果如下：
```shell
$ go build q2.go
msg.Name undefined (type interface {} is interface with no methods) 
```

有两个问题：
问题一：interface{}是一个没有声明任何方法的接口。
问题二：Name是一个属性，而不是方法，interface{}类型的变量无法调用属性。

改成下面的函数，就能正常编译运行
```go
func hello(v interface{}) {
    switch msg := v.(type) {
    // case *student, student: // case 有多个类型，则 msg 变量的类型是 interface 
    case student: // case 只有一个类型，则 msg 变量的类型是 student 
        fmt.Println(msg.Name)
    }
}
```

> 这个题目过于细节，并不适合作为面试题目
> 考察的应该 Interface conversions and type assertions 相关知识。
> 网上有相关讨论：[https://github.com/golang/go/issues/12772](https://github.com/golang/go/issues/12772)
> 参考 [effective go](https://golang.org/doc/effective_go.html#conversions) [spec](https://golang.org/ref/spec#Conversions) [go101](https://go101.org/article/type-system-overview.html)
> type conversion 用于具体的 type 转换到另外一个具体的 type ，比如 `string([]byte{'a'}) `
> Interface conversions ( type assertions / type switch ) 用于 interface 转换到 其他类型 (type 或 interface) ，`

> type conversion
```go
MyString("foo" + "bar")  // "foobar" of type MyString
string([]byte{'a'})      // not a constant: []byte{'a'} is not a constant
(*int)(nil)              // not a constant: nil is not a constant, *int is not a boolean, numeric, or string type
```

> type assertion
```go
str, ok := value.(string)
if ok {
    fmt.Printf("string value is: %q\n", str)
} else {
    fmt.Printf("value is not a string\n")
}
```

> type switch
```go
var value interface{} // Value provided by caller.
switch str := value.(type) {
case string:
    return str
case Stringer:
    return str.String()
}
```


## 3. 请找出下面代码的问题所在。
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}
```


运行结果如下：
```shell
panic: sendon closed channel
ok

goroutine 5 [running]:
main.main.func1(0xc420098000)
```


解析：出现上面错误的原因是因为提前关闭通道所致。

相关问题：
- 1.有没有更好的方法，替换 `time.Sleep()`? sync.WaitGroup
- 2.多个goroutine同时写一个int型变量会有什么问题？同时对int变量执行 `i++`自增操作有什么问题？ 原子操作

正确代码如下：
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个缓冲通道
	ch := make(chan int, 1000)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			a, ok := <-ch

			if !ok {
				fmt.Println("close")
				close(ch)
				return
			}
			fmt.Println("a: ", a)
		}
	}()

	fmt.Println("ok")
	time.Sleep(time.Second)
}
```

运行结果如下：
```shell
ok
a:  0
a:  1
a:  2
a:  3
a:  4
a:  5
a:  6
a:  7
a:  8
a:  9
```



## 4. 下面的程序运行后为什么会爆异常。
```go
package main

import (
	"fmt"
	"time"
)

type Project struct{}

func (p *Project) deferError() {
	if err := recover(); err != nil {
		fmt.Println("recover: ", err)
	}
}

func (p *Project) exec(msgchan chan interface{}) {
	for msg := range msgchan {
		m := msg.(int)
		fmt.Println("msg: ", m)
	}
}

func (p *Project) run(msgchan chan interface{}) {
	for {
		defer p.deferError()
		go p.exec(msgchan)
		time.Sleep(time.Second * 2)
	}
}

func (p *Project) Main() {
	a := make(chan interface{}, 100)
	go p.run(a)
	go func() {
		for {
			a <- "1"
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 100)
}

func main() {
	p := new(Project)
	p.Main()
}
```

运行结果如下：
```shell
panic: interface conversion: interface {} isstring, not int

goroutine 17 [running]:
main.(*Project).exec(0x1157c08, 0xc420068060)
```


出现异常的原因是因为写入到管道的数据类型为string,而m := msg.(int)这句代码里面却使用了int，修改方法，将int修改为string即可。

相关问题：
- 1.类型断言与类型转换的区别？ 处理 type 时使用 type conversion ，处理 interface 时 使用 type assertion

[https://stackoverflow.com/questions/20494229/what-is-the-difference-between-type-conversion-and-type-assertion](https://stackoverflow.com/questions/20494229/what-is-the-difference-between-type-conversion-and-type-assertion)

完整正确代码如下：
```go
package main

import (
	"fmt"
	"time"
)

type Project struct{}

func (p *Project) deferError() {
	if err := recover(); err != nil {
		fmt.Println("recover: ", err)
	}
}

func (p *Project) exec(msgchan chan interface{}) {
	for msg := range msgchan {
		m := msg.(string)
		fmt.Println("msg: ", m)
	}
}

func (p *Project) run(msgchan chan interface{}) {
	for {
		defer p.deferError()
		go p.exec(msgchan)
		time.Sleep(time.Second * 2)
	}
}

func (p *Project) Main() {
	a := make(chan interface{}, 100)
	go p.run(a)
	go func() {
		for {
			a <- "1"
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 100)
}

func main() {
	p := new(Project)
	p.Main()
}
```

运行结果如下：
```shell
msg:1
msg:1
msg:1
msg:1
msg:1
msg:1
msg:1
msg:1
```


## 5. 请说出下面代码，执行时为什么会报错
```go
package main

import ()

type Student struct {
	name string
}

func main() {
	m := map[string]Student{"people": {"liyuechun"}}
	m["people"].name = "wuyanzu"
}
```

答案：报错的原因是因为不能修改字典中value为结构体的属性值。


代码作如下修改方可运行：
```go
package main

import "fmt"

type Student struct {
	name string
}

func main() {
	m := map[string]Student{"people": {"liyuechun"}}
	fmt.Println(m)
	fmt.Println(m["people"])

	// 不能修改字典中结构体属性的值
	//m["people"].name = "wuyanzu"

	var s Student = m["people"] //深拷贝
	s.name = "xietingfeng"
	fmt.Println(m)
	fmt.Println(s)
}
```

运行结果如下：
```shell
map[people:{liyuechun}]
{liyuechun}
map[people:{liyuechun}]
{wuyanzu}
```

相关问题：
- 1. Golang 中值与引用的区别？哪些是引用类型？
- 2. 哪些是指针类型？ TODO

[https://github.com/go101/go101/wiki/About-the-terminology-%22reference-type%22-in-Go](https://github.com/go101/go101/wiki/About-the-terminology-%22reference-type%22-in-Go)

GO 中没有引用类型， 只有值和指针。 引用类型没有官方的业格定义，但官方 blogs 中有出现几次。


## 6．tcp与udp区别，udp优点，适用场景

### 6.1 一台 tcp server 主机上运行一个进程 listen 一个 8080 port ,理论最多支持多少并发连接

## 7. Slice与数组区别，Slice底层结构

## 8. git rebase 和 git merge 区别
<!--stackedit_data:
eyJoaXN0b3J5IjpbLTIxMTQyNzcyNDIsMTU5NzI1MzQwNSwtNj
Q5MDcyMDA4LC02OTE4NDQ3OSw0NjQ0MjY5MzMsMTExODQ3Mjkz
MywtMTI2ODkyOTY3NCwtMTEwOTg3MTg5Niw0NjQ1NTYzNDksMT
MzNzU0MzY5N119
-->
