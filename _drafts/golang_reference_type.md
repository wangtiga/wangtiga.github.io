
# About the terminology "reference type" in Go[^reference_type_in_go]

Although the terminology  _**reference type**_  [has been removed from Go specification since Apr 3rd, 2013](https://github.com/golang/go/commit/b34f0551387fcf043d65cd7d96a0214956578f94)  (with the commit message:  `spec: Go has no 'reference types'`), the terminology is still popularly used in Go community. One reason is many people think this terminology is necessary, for it will make many explanations easy and simple. However, I never saw an accurate definition for the terminology  _**reference type**_  in any official documents and unofficial Go articles. And in fact, the confusions brought by this terminology are more than the explanation conveniences it brings. The following of this article will list all kinds of reference type definitions and point out the inaccuracies in them.

(Firstly, I would present one of my personal opinion: I very dislike to view  _**reference type**_  and  _**value type**_  as two opposite concepts. The terminology  _**value**_  means an instance of a type. In my opinion, any type is a value type. So, in the following content, I will use the  _**non-reference type**_  to represent the  _**value type**_  concept used in many other articles on Internet.)

At present (Aug 26th, 2018), the only official Go document which uses the  _**reference type**_  terminology in a clear manner is this blog article:  [Go maps in action](https://blog.golang.org/go-maps-in-action). In this article, it says

> Map types are reference types, like pointers or slices, and so the value of m above is nil; ...

This article clearly states that map, pointer and slice types are reference types. And it seems also hints that any type whose zero value is  `nil`  is also a reference type. So this blog article makes two possible reference type definitions:

> 译： zero value 中 golang 中是一个专有名词。具体批哪些？ effective golang 中也有 zero value 的描述

1.  map, pointer and slice types are reference types in Go.
2.  map, pointer, slice, channel, function and interface types are reference types.

There are two defects for the first definition. The first defect is there are no reasons why map types are viewed as reference types, but channel types aren't. The second defect is that slice types also present some non-reference type behaviors. For example:

第一个结论的问题是，为什么 map 是 reference type ，但 channel 不是？
第二个结论的问题是，slice type 有时也会表现出 非 reference type 的行为。

```go
package main

import (
	"fmt"
	"reflect"
)

func modifySlice(s []int) {
	reflect.ValueOf(&s).Elem().SetLen(4)
	reflect.ValueOf(&s).Elem().SetCap(4)
	fmt.Println("b>", len(s), cap(s)) // b> 4 4
}

func main() {
	s := make([]int, 3, 5)
	fmt.Println("a>", len(s), cap(s)) // a> 3 5
	modifySlice(s)
	fmt.Println("c>", len(s), cap(s)) // c> 3 5
}
```

The above example shows that, at least for some specified circumstances, slice types will present some non-reference type behaviors.

There are three defects in the above possible second definition:

1.  the first defect is one is the same as the second defect described above for the first possible definition.
2.  the second defect is that there are no ways to verify function types are reference types or not, without using the unsafe mechanism, for function values are immutable internally. So it is okay to view function types as both reference types and non-reference types.
3.  the third defect is that viewing interface types as reference types is totally wrong. The official Go FAQ  [states it clearly](https://golang.org/doc/faq#pass_by_value)

> ... Copying an interface value makes a copy of the thing stored in the interface value. If the interface value holds a struct, copying the interface value makes a copy of the struct. ...

So interface types are non-reference types, at least in theory, though the standard Go compiler (gc) and gccgo both implement interface types with a reference type manner, for code execution optimization reason. On the other hand, same as function types, the dynamic value of any interface value is immutable internally, so without using the unsafe mechanism, there are no ways to verify whether interface types are non-reference types or reference types. In fact, if interface types are viewed as reference types, then string types should be too.

Another place seeming to use the  _**reference type**_  terminology in official documents is  [this FAQ item](https://golang.org/doc/faq#references)  in the official Go FAQ:

> Why are maps, slices, and channels references while arrays are values?
> 
> There's a lot of history on that topic. Early on, maps and channels were syntactically pointers and it was impossible to declare or use a non-pointer instance. Also, we struggled with how arrays should work. Eventually we decided that the strict separation of pointers and values made the language harder to use. Changing these types to act as references to the associated, shared data structures resolved these issues. This change added some regrettable complexity to the language but had a large effect on usability: Go became a more productive, comfortable language when it was introduced.

> 为什么 maps, slices, channel 是引用类型，但  array 是值类型
> 
> 这是由于一些历史原因造成。早期 map 和 channel 必须使用指针的语法来使用。
> 同时我们也在考虑 array 应该怎么工作。
> 最终，我们认为严格地区分 pointer 和 value 会使语言变得很难使用。
> 如果把这些类型当做 reference 到共享数据结构的来使用，会方便一些。
> 这样会使语言变得复杂，但更容易使用: Go应该是一个极具生产力，易于使用的语言。

IMHO, this FAQ item is very misleading, in several ways.

-   First, it uses terminology  _**reference**_  as an opposite of  _**value**_, which is not good. (OK, this is my personal opinion.)
-   Second, what are exactly the terminology  _**reference**_  used here? Reference types? Reference values? Or others? Here, we will think it as  _**reference type**_.

When the terminology  _**reference**_  in the above FAQ item is interpreted as  _**reference type**_, we can think the FAQ item defines  _**map, slice and channel types are reference types**_. In fact, this definition is a comparative good one among all definitions. It just has one inaccuracy, which has been mentioned above, slice types may present some non-reference type behaviors sometimes.

(Note, the definitions from the two different official Go articles mentioned above are different.)

Another popular definition of  _**reference type**_  is  _**reference types are the types whose values may reference other values**_. By this definition, an array type with a reference element type and a struct type with fields of reference types, along with map/slice/channel/pointer/interface types, are verifiable reference types. Interface types are included here is for the dynamic value of an interface value may reference other values. This definition is not bad if it mentions the following special notes

1.  slice and reference array/struct types may present some non-reference type behaviors sometimes.
2.  an interface value with a non-interface dynamic type presents non-reference type behaviors.

The two special notes make the just described definition lose the advantage of explanation convenience.

As it is so hard to make a both accurate and useful definition for  _**reference type**_, to void the confusions it brings, I recommend not to use this terminology in any Go documents and articles at all. In fact, I think this terminology is not very necessary. By knowing that reference relations are built through pointers and  [some possible internal structures of all kinds of built-in kinds of types](https://go101.org/article/value-part.html), there would be less confusions in using Go without using the  _**reference type**_  terminology.

_[update]: On the other hand, I do think  **non-references type**  is a clear concept. All basic types can be viewed as non-references types. Array types with non-references element type are also non-references types. Struct types with app fields of no-nreferences types are also non-reference types. Personally, I tend to think function types are also non-reference types._



[^reference_type_in_go]: [reference type" in Go](https://github.com/go101/go101/wiki/About-the-terminology-%22reference-type%22-in-Go)


<!--stackedit_data:
eyJoaXN0b3J5IjpbNDYyNjAyODk2LDE4OTY5NTY2MjUsLTk2Mj
g2MTY0MywxMTA5MDMxNDgxXX0=
-->
