---
title: golang 的文件锁操作
---
这篇文章给大家介绍一下 golang 的文件锁。我们在使用 golang 开发程序的时候，经常会出现多个 goroutine 操作同一个文件（或目录）的时候，如果不加锁，很容易导致文件中的数据混乱，于是，Flock 应运而生。

Flock 是对于整个文件的建议性锁（不强求 goroutine 遵守），如果一个 goroutine 在文件上获取了锁，那么其他 goroutine 是可以知道的。默认情况下，当一个 goroutine 将文件锁住，另外一个 goroutine 可以直接操作被锁住的文件，原因在于 Flock 只是用于检测文件是否被加锁，针对文件已经被加锁，另一个 goroutine 写入数据的情况，内核不会阻止这个 goroutine 的写入操作，也就是建议性锁的内核处理策略。

## 函数

```go
import "syscall"

func Flock(fd int, how int) (err error)
```

Flock 位于 syscall 包中，fd 参数指代文件描述符，how 参数指代锁的操作类型。

how 主要的参数类型：

* LOCK_SH，共享锁，多个进程可以使用同一把锁，常被用作读共享锁；
* LOCK_EX，排他锁，同时只允许一个进程使用，常被用作写锁；
* LOCK_NB，遇到锁的表现，当采用排他锁的时候，默认 goroutine 会被阻塞等待锁被释放，采用 LOCK_NB 参数，可以让 goroutine 返回 Error;
* LOCK_UN，释放锁；

## 示例

下面的例子来自于 NSQ，位于 `nsq/internal/dirlock`，用于实现对目录的加锁

```go
// +build !windows

package dirlock

import (
	"fmt"
	"os"
	"syscall"
)

// 定义一个 DirLock 的struct
type DirLock struct {
	dir string    // 目录路径，例如 /home/XXX/go/src
	f   *os.File  // 文件描述符
}

// 新建一个 DirLock
func New(dir string) *DirLock {
	return &DirLock{
		dir: dir,
	}
}

// 加锁操作
func (l *DirLock) Lock() error {
	f, err := os.Open(l.dir) // 获取文件描述符
	if err != nil {
		return err
	}
	l.f = f
	err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB) // 加上排他锁，当遇到文件加锁的情况直接返回 Error
	if err != nil {
		return fmt.Errorf("cannot flock directory %s - %s", l.dir, err)
	}
	return nil
}

// 解锁操作
func (l *DirLock) Unlock() error {
	defer l.f.Close() // close 掉文件描述符
	return syscall.Flock(int(l.f.Fd()), syscall.LOCK_UN) // 释放 Flock 文件锁
}
```

## 总结

1. Flock 是建议性的锁，使用的时候需要指定 `how` 参数，否则容易出现多个 goroutine 共用文件的问题
2. `how` 参数指定 `LOCK_NB` 之后，goroutine 遇到已加锁的 Flock，不会阻塞，而是直接返回错误
