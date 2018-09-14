## 2018-09-14

来源：《Go 夜读》微信群

### VSCode 如何代码自动补全和自动导入包

VSCode 必须安装以下插件：

首先你必须安装 Golang 插件，然后再给 Go 安装工具包。

在 VS Code 中，使用快捷键：`command+shift+P`，然后键入：`go:install/update tools`，将所有 16 个插件都勾选上，然后点击 OK 即开始安装。

```
Installing 16 tools at /Users/maiyang/develop/goworkspace//bin
  gocode
  gopkgs
  go-outline
  go-symbols
  guru
  gorename
  dlv
  godef
  godoc
  goreturns
  golint
  gotests
  gomodifytags
  impl
  fillstruct
  goplay

Installing github.com/mdempsky/gocode SUCCEEDED
Installing github.com/uudashr/gopkgs/cmd/gopkgs SUCCEEDED
Installing github.com/ramya-rao-a/go-outline SUCCEEDED
Installing github.com/acroca/go-symbols SUCCEEDED
Installing golang.org/x/tools/cmd/guru SUCCEEDED
Installing golang.org/x/tools/cmd/gorename SUCCEEDED
Installing github.com/derekparker/delve/cmd/dlv SUCCEEDED
Installing github.com/rogpeppe/godef SUCCEEDED
Installing golang.org/x/tools/cmd/godoc SUCCEEDED
Installing github.com/sqs/goreturns SUCCEEDED
Installing github.com/golang/lint/golint SUCCEEDED
Installing github.com/cweill/gotests/... SUCCEEDED
Installing github.com/fatih/gomodifytags SUCCEEDED
Installing github.com/josharian/impl SUCCEEDED
Installing github.com/davidrjenni/reftools/cmd/fillstruct SUCCEEDED
Installing github.com/haya14busa/goplay/cmd/goplay SUCCEEDED

All tools successfully installed. You're ready to Go :).
```

然后在 Preferences -> Setting 然后输入 go，然后选择 `setting.json`，填入以下配置（自动完成未导入的包）：

```json
  "go.autocompleteUnimportedPackages": true,
```

特别注意：

VS Code 需要安装好代理，才能够正常安装，否则可能无法安装成功。

## 其他

1. 当我们在使用 import 功能的时候，如果无法通过 lint 检查，则不会执行自动 import。
2. 如果你需要自动 import 的前提是你必须把要导入的包的函数写完整。

