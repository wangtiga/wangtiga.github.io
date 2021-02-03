---
layout: page
title: Preview
---


{% assign posts_by_year = site.previews | group_by_exp: "post", "post.date | date: '%Y' " %}
{% for group in posts_by_year %}

<h3 class="group-name">{{ group.name }}</h3>

<ul>
  {% for post in group.items %}
    <li>
			<div class="group-date">{{ post.date | date: "%b %-d" }}</div> 
		  <a href="{{ post.url | relative_url }}">{{ post.title }}</a>
		</li>
  {% endfor %}
</ul>
{% endfor %}
