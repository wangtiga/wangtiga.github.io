

```shell
#gem list --local  | grep jekyll
#gem uninstall jekyll --version 3.8.6
#gem install jekyll --version 3.8.6
#bundle exec jekyll serve

# Quick-start Instructions https://jekyllrb.com/
gem install bundler jekyll
jekyll new my-awesome-site
cd my-awesome-site
bundle exec jekyll serve
# => Now browse to http://localhost:4000
```

以下是个人的浅显理解，没有验证过真伪

- Gem
  
  gem 可理解为包管理器，是用 ruby 语言实现的

- Gemfile 

  记录当前项目依赖的第三方库，及相关版本。

- Gemfile.lock 

  某些命令后，当前使用的第三方库的版本会记录在这个文件，下次再执行相关命令时，不会自动更新版本。除非删除这个文件

- Bundler

  bundle 是包含所有相关三方库的命令，可把它理解成一个运行环境

- jekyll 。

  jekyll 也一个 ruby 程序，主要实现将 markdown 翻译成 html 语言，生成 blog 所需要的网页文件

- `_config.yml`

  记录当前 blogs 的环境配置，比如博客的名称，联系人的邮箱等。

  TODO 也会记录当前 blogs 使用的第三方库，可能是给 github pages 后台服务使用的吧？
  但是为什么要在 Gemfile 与 `_config.yml` 两个地方分别记录第三方库呢？


https://jekyllrb.com/docs/ruby-101/

https://stackoverflow.com/questions/58195772/how-can-fix-the-jekyll-installation-error-in-termux

https://wiki.termux.com/wiki/List_of_Ruby_packages_known_to_work

https://wiki.termux.com/wiki/Jekyll

