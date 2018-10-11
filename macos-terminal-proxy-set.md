Macos 终端设置代理
===============

前提是需要有个http代理 例如端口号为 1081

下面的命令加入到 ` ~/.bash_profile`（如果使用系统终端) 或者 `~/.zshrc`（如果使用zsh)，更改完成之后请在当前终端执行 `source ~/.bash_profile` or `source ~/.zshrc`

```
alias setproxy="export http_proxy=http://127.0.0.1:1081 && export https_proxy=http://127.0.0.1:1081 && curl -i http://ip.cn"
alias unsetproxy="unset http_proxy && unset https_proxy && curl -i http://ip.cn"
```
上面的方式来自群友 @Amor 


因为我的代理设置了pac，所以使用 ip.cn还是被转发到国内， 测试就不准了，于是使用 `google.com` 来作为测试url， 使用方式同上
```
alias setproxy="export http_proxy=http://127.0.0.1:1081 && export https_proxy=http://127.0.0.1:1081 && curl -i http://google.com"
alias unsetproxy="unset http_proxy && unset https_proxy && curl -i http://google.com"
```


使用
```
▶ setproxy
HTTP/1.1 301 Moved Permanently
Content-Length: 219
Date: Thu, 11 Oct 2018 01:58:28 GMT
Expires: Sat, 10 Nov 2018 01:58:28 GMT
Cache-Control: public
Location: http://www.google.com/
Content-Type: text/html; charset=UTF-8
Server: gws
X-XSS-Protection: 1; mode=block
X-Frame-Options: SAMEORIGIN
Age: 171
Connection: keep-alive

<HTML><HEAD><meta http-equiv="content-type" content="text/html;charset=utf-8">
<TITLE>301 Moved</TITLE></HEAD><BODY>
<H1>301 Moved</H1>
The document has moved
<A HREF="http://www.google.com/">here</A>.
</BODY></HTML>

~
▶ unsetproxy
^C
```
