---
title: 2018-09-14 VSCode 如何代码自动补全和自动导入包
---
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

修改默认配置的方法：

>在 Preferences -> Setting 然后输入 go，然后选择 `setting.json`，填入你想要修改的配置


- 自动完成未导入的包。

```json
  "go.autocompleteUnimportedPackages": true,
```

- VSCode 的一些插件需要配置代理，才能够正常安装。

```json
  "http.proxy": "192.168.0.100:1087",
```

- 如果你遇到使用标准包可以出现代码提示，但是使用自己的包或者第三方库无法出现代码提示，你可以查看一下你的配置项。

```json
  "go.inferGopath": true,
```

- 如果引用的包使用了 ( . "aa.com/text") 那这个text包下的函数也无法跳转进去，这是为什么？

修改 `"go.docsTool"` 为 `gogetdoc`，默认是 `godoc`。

```json
  "go.docsTool": "gogetdoc",
```

## 其他

1. 当我们在使用 import 功能的时候，如果无法通过 lint 检查，则不会执行自动 import。
2. 如果你需要自动 import 的前提是你必须把要导入的包的函数写完整。

附带我的 `settings.json`

```json
{
  "go.goroot": "",
  "go.gopath": "",
  "go.inferGopath": true,
  "go.autocompleteUnimportedPackages": true,
  "go.gocodePackageLookupMode": "go",
  "go.gotoSymbol.includeImports": true,
  "go.useCodeSnippetsOnFunctionSuggest": true,
  "go.useCodeSnippetsOnFunctionSuggestWithoutType": true,
  "go.docsTool": "gogetdoc",
}
```

## 参考资料

1. [GOPATH in the VS Code Go extension](https://github.com/Microsoft/vscode-go/wiki/GOPATH-in-the-VS-Code-Go-extension)
2. [VSCode Golang 开发配置之代码提示](https://www.cnblogs.com/Dennis-mi/p/8280552.html)
3. [Use gogetdoc instead of godef and godoc #622](https://github.com/Microsoft/vscode-go/pull/622)
