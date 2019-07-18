---
desc: Go 夜读之 sync/errgroup 源码阅读
title: 第 51 期 Go 夜读之 sync/errgroup 源码阅读
date: 2019-07-18T21:10:00+08:00
author: mai
---

## golang.org/x/sync/errgroup

errgroup 唯一的坑是for循环里千万别忘了 i, x := i, x，以前用 waitgroup 的时候都是 go func 手动给闭包传参解决这个问题的，errgroup 的.Go没法这么干，犯了好几次错才改过来"

## 观看视频

{{< youtube id="CQOZtzmgLvw" >}}
