---
title: 2018-09-19 微信讨论
---
来源: Wechat discuss

时间：2018-09-19

## Producer to Consumer via channel
这段code有问题，可能程序跑完了，buffer中还存在数据，就丢失了。

```golang
package main

import (
	"fmt"
)

const maxBufSize = 3
const numToProduce = 1000

var finishedProducing = make(chan bool)
var finishedConsuming = make(chan bool)

var messageBuffer = make(chan int, maxBufSize)

func produce() {
	for i := 0; i < numToProduce; i++ {
		messageBuffer <- i
	}

	finishedProducing <- true
}

func consume() {
	for {
		select {
		case message := <-messageBuffer:
			fmt.Println(message)
		case <-finishedProducing:
			finishedConsuming <- true
			return
		}
	}
}

func main() {
	go produce()
	go consume()
	<-finishedConsuming

	fmt.Println("All go routines ended")
}

```

## 原因
buffer中的消息可能没有被处理


## 解决办法

应该用for range从channel 里取，produce 那里直接关闭messageChannel就行了。 （AMan）

```golang
package main

import (
	"fmt"
)

const maxBufSize = 3
const numToProduce = 1000

// var finishedProducing = make(chan bool)
var finishedConsuming = make(chan bool)

var messageBuffer = make(chan int, maxBufSize)

func produce() {
	for i := 0; i < numToProduce; i++ {
		messageBuffer <- i
	}

	close(messageBuffer)
}

func consume() {
	for message := range messageBuffer {
		fmt.Println(message)
	}
	finishedConsuming <- true

}

func main() {
	go produce()
	go consume()
	<-finishedConsuming

	fmt.Println("All go routines ended")
}

```

## 提升

可以使用 `make(chan struct{})` 代替 `make(chan bool)`
