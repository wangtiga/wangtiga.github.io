---
layout: post
title:  markdown template
date:   1990-01-01 13:00:00 +0800
tags:   todo

mermaid: true
mathjax: true
---

The article introduces the features of the Jekyll theme support.

## Paragraph

In a paragraph, different fonts may appear, including but not limited to:

1. *Italics*

2. **Bold**

3. ~~Delete~~

4. <u>Underline</u>

## Link

If you have any questions, just [Google][google-link] it.

[google-link]: http://www.google.com/

## Footnote

Knowledge is power. [^professordeng]。

[^professordeng]: Professordeng said.

## Quote

> Nobody knows politics better than me. --Trump

## Image

Insert a image by `![image name](image url)`, like:

![picture]({{ site.img }}/big.jpg)

## Table

Normal Table:

|       | mathematics | English | Chinese | Politics |
| ----- | ----------- | ------- | ------- | -------- |
| David | 80          | 80      | 50      | 100      |
| James | 70          | 80      | 90      | 100      |
| Me    | 100         | 90      | 100     | 100      |

Super Wide Table:

|       | mathematics | English | Chinese | Politics | Japanese | python | basketball | javascript |
| ----- | ----------- | ------- | ------- | -------- | -------- | ------ | ---------- | ---------- |
| David | 80          | 80      | 50      | 100      | 98       | 100    | 99         | 1          |
| James | 70          | 80      | 90      | 100      | 12       | 90     | 88         | 2          |
| Me    | 100         | 90      | 100     | 100      | 120      | 50     | 77         | 3          |

## Code

```c
int main() {
	printf("hello world!");
	return 1;
}
```

## Todo List

Unordered list

- [x] This is a complete item.
- [ ] This is an incomplete item.

Ordered list

1. [x] This is a complete item.
2. [ ] This is an incomplete item.

## Emoji

The Jekyll theme support [Emoji](https://emojipedia.org/).

- 🙌 👐 👌 ☝ : Nobody know more about technology than me.
	
- Also you can use `:smile:` to show :smile:.

## Video

`video` Label：

<video src="https://cdn-video.xinpianchang.com/5b7fc02a84108.mp4" controls controlsList="nodownload"></video>

Bilibili `iframe` Label：

<iframe class="video" src="//player.bilibili.com/player.html?bvid=BV1ki4y1b7ge&page=1&high_quality=1&danmaku=0" allowfullscreen> </iframe>

YouTube `iframe` Label:

<iframe class="video" src="https://www.youtube.com/embed/-wFsYY71wyk" allowfullscreen></iframe>

## Mentions

You can mention somebody by `@username`, just like @professordeng.

## Avatar

Maybe you need [jekyll-avatar](https://github.com/benbalter/jekyll-avatar/) to show an avatar list, just like:

{% avatar professordeng %}
{% avatar ericclose %}

If you want to transform avatar to a link, you can use `[avatar link](url)` format, try click the avatar:

[{% avatar professordeng %}](https://github.com/professordeng)
[{% avatar ericclose %}](https://github.com/ericclose)

## Formula

The relationship between $c$ and $r$:

$$
c=2 \pi r
$$

Formula is supported by [MathJax](https://www.mathjax.org/)

## Flowchart

<div class="mermaid">
graph LR
    A --- B
    B-->C[fa:fa-ban forbidden]
    B-->D(fa:fa-spinner);
</div>

Flowchart is supported by [mermaid](https://mermaid-js.github.io/mermaid/#/)

---
