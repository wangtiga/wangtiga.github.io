---
layout: post
title:  "url 格式"
date:   2020-02-08 18:00:00 +0800
tags:   tech
---


* category
{:toc}


## url 

统一资源定位符（英語：Uniform Resource Locator，缩写：URL；或称统一資源定位器、定位地址、URL地址[1]，俗称网页地址或简称网址）

```txt
http://example.domain.com/path/subpath?key1=val1&key2=val2
```

- path 编码
主要为了转义 `/` 符号

- key/value 编码
主要为了转义 `?` `=` `&` 等符号

- `#` 

- reference net/url/url.go
```go
// QueryEscape escapes the string so it can be safely placed
// inside a URL query.
func QueryEscape(s string) string {
        return escape(s, encodeQueryComponent)
}

// PathEscape escapes the string so it can be safely placed inside a URL path segment,
// replacing special characters (including /) with %XX sequences as needed.
func PathEscape(s string) string {
        return escape(s, encodePathSegment)
}

// Return true if the specified character should be escaped when
// appearing in a URL string, according to RFC 3986.
//
// Please be informed that for now shouldEscape does not check all
// reserved characters correctly. See golang.org/issue/5684.
func shouldEscape(c byte, mode encoding) bool {
// ...
}
```



### 几个特殊的URL地址

```txt
# error
http://example.com/video#abc.mp4

# error
http://example.com/video#abc.mp4?

```


[golang doc](https://godoc.org/net/url#URL)
[wikipedia](https://zh.wikipedia.org/wiki/%E7%BB%9F%E4%B8%80%E8%B5%84%E6%BA%90%E5%AE%9A%E4%BD%8D%E7%AC%A6)

