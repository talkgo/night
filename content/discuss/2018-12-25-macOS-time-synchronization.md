---
title: 2018-12-25 MacOS Time Synchronization
date: 2018-12-25T16:36:45+08:00
---
来源: Wechat discuss

## MacOS Time Synchronization

> 今天在启动 Geth (以太坊私链节点)的时候，出现了以下的错误。在比较后发现本机时间落后了北京时间近一分钟。

![](/images/2018-12-25-macOS-time-synchronization-01.jpg)


> 在搜索解决方案的时候，发现在 macOS Mojave 里面 `ntpdate` 已经废弃了。
> 改用了`sntp`，在Terminal输入以下的命令，即可同步。


```bash
sudo sntp -sS time.asia.apple.com
```

### 系统设置建议

![](/images/2018-12-25-macOS-time-synchronization-02.jpg)

- 1、自动设置日期与时间
- 2、时区不要自动设定，选择中国北京，因为不同地区的时间会有误差。


## 参考资料

1. https://apple.stackexchange.com/questions/117864/how-can-i-tell-if-my-mac-is-keeping-the-clock-updated-properly