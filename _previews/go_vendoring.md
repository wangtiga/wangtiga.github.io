




- 环境
  1. go version go1.16.3 linux/amd64
  2. go.mod `go 1.13`

  在 vim-go 中用 Ctrl + ] 查看三方库时，会跳转到 $GOPATH/pkg/mod 目录

- 更新  go.mod 版本 为 `go 1.13` 后

  在 vim-go 中用 Ctrl + ] 查看三方库时，会提示 `vim-go: : packages.Load error`

- 执行 `:messages` 查看详细报错信息，提示 `go: inconsistent vendoring `

```vim
vim-go: [definition] FAIL
vim-go: err: exit status 1: stderr: go: inconsistent vendoring in /home/tiga/src/testserver:
vim-go:         github.com/stretchr/testify@v1.6.1: is explicitly required in go.mod, but not marked as explicit in vendor/modules.txt
vim-go:         go.uber.org/zap@v1.16.0: is explicitly required in go.mod, but not marked as explicit in vendor/modules.txt
vim-go: 
vim-go:         To ignore the vendor directory, use -mod=readonly or -mod=mod.
vim-go:         To sync the vendor directory, run:
vim-go:         ^Igo mod vendor
vim-go: : packages.Load error
请按 ENTER 或其它命令继续
```

- 解决 inconsistent vendoring [x/tools/gopls: inconsistent vendoring issue #37734](https://github.com/golang/go/issues/37734)

  先执行 `go mod tidy` ，再执行  `go mod vendor`

  1. This error is reported when automatic vendoring is enabled and go.mod is not consistent with vendor/modules.txt. 
  2. Automatic vendoring is enabled if the go.mod file has go version 1.14 or later and a vendor directory is present. 
  3. go mod tidy ensures that every necessary requirement appears in go.mod. 
  4. go mod vendor copies imported packages into vendor/ and updates vendor/modules.txt.


- 再次尝试在 vim-go 中用 Ctrl + ] 查看三方库，顺利跳转到 vendor 目录中的代码，而非 $GOPATH/pkg/mod 目录。


- git diff vendor/modules.txt 查看差异，保是增加了 `## explicit` 行，可能是 go.mod 1.14 版本变化新增的语法。

相关文章 [Modules Part 06: Vendoring](https://www.ardanlabs.com/blog/2020/04/modules-06-vendoring.html)

```diff
diff --git a/vendor/modules.txt b/vendor/modules.txt
index 1be2340..a1ff23a 401234
--- a/vendor/modules.txt
+++ b/vendor/modules.txt
@@ -32,19 +35,22 @@ github.com/eapache/queue
-# github.com/stretchr/testify v1.5.1
+# github.com/stretchr/testify v1.6.1
+## explicit
 github.com/stretchr/testify/assert
 github.com/stretchr/testify/require
-# go.uber.org/atomic v1.7.0
+# go.uber.org/atomic v1.6.0
 go.uber.org/atomic
+## explicit
 go.uber.org/zap
```


