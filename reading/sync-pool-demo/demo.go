package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

// 临时对象池
var p = sync.Pool{
	New: func() interface{} {
		buffer := make([]byte, 256)
		return &buffer
	},
}

//wg 是一个指针类型,必须是一个内存地址
func readContent(wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()

	byteSlice := p.Get().(*[]byte) //类型断言

	numBytesReadAtLeast, err := io.ReadFull(resp.Body, *byteSlice)
	if err != nil {
		// handle error
	}

	p.Put(byteSlice)

	log.Printf("Number of bytes read: %d\n", numBytesReadAtLeast)
	fmt.Println(string((*byteSlice)[:256]))
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go readContent(&wg)
	}

	wg.Wait()
	fmt.Println("end...")
}
