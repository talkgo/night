package main

import (
	"log"
	"runtime"
	"sync"
)

func main() {
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}
	a := p.Get().(int)
	p.Put(1)
	b := p.Get().(int)
	log.Println(a, b)
	p.Put(3)
	p.Put(4)
	p.Put(5)
	log.Println(p.Get()) //返回 3 4 5中的任意一个。
	//主动调用GC  pool中对象会被清理掉
	runtime.GC()
	p.Put(2)
	c := p.Get().(int)
	log.Println(c)
}
