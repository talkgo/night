---
title: "how to testing"
date: 2018-12-25T00:00:00+08:00
author: xpzouying@gmail.com
---


# HOW TO TESTING

原文/源码参考：

- [how_to_test](https://github.com/xpzouying/learning_golang/tree/master/how_to_test)

作者：xpzouying@gmail.com

---

测试的作用：

- 验证代码是否符合预期
- 资源竞争检查：race detect
- 调优：profiling：memory/cpu



## 原始代码

代码功能：访客记次数。

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

var counter = map[string]int{}

func handleHello(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	counter[name]++

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<h1 style='color: " + r.FormValue("color") +
		"'>Welcome!</h1> <p>Name: " + name + "</p> <p>Count: " + fmt.Sprint(counter[name]) + "</p>"))
}

func main() {
	http.HandleFunc("/hello", handleHello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```



运行：

```bash
go run main.go
```



浏览器访问：

![image-20181222170909680](/static/images/image-20181222170909680-5469749.png)

本地日志记录：

![image-20181222171507017](/static/images/image-20181222171507017.jpg)



### 测试规范

1. 运行测试：

   1. 测试：`go test`
   2. 压力测试：`go test -bench`
   3. 测试覆盖：`go test -cover`

2. 测试规范：

   1. 测试函数示例

      ```go
      // go test or go test -v
      func TestXxx(*testing.T)
      
      // go test -bench
      func BenchmarkXxx(*testing.B)
      ```

      Xxx不能以小写字母开头。

   2. 测试文件规范：文件名以`_test.go`结尾。

   3. 在测试函数里面使用：Error，Fail或者相关的函数标示相关错误。

3. 例子：

   1. 单元测试：

      ```go
      func TestTimeConsuming(t *testing.T) {
          if testing.Short() {
              t.Skip("skipping test in short mode.")
          }
          ...
      }
      ```

   2. 压力测试：

      ```go
      func BenchmarkHello(b *testing.B) {
          for i := 0; i < b.N; i++ {
              fmt.Sprintf("hello")
          }
      }
      ```

   3. Examples：

      ```go
      func ExampleHello() {
          fmt.Println("hello")
          // Output: hello
      }
      
      func ExampleSalutations() {
          fmt.Println("hello, and")
          fmt.Println("goodbye")
          // Output:
          // hello, and
          // goodbye
      }
      ```



### 测试用例



**运行测试**

使用`go test`运行测试。

```bash
➜  how_to_test git:(how_to_test) ✗ go test
?       _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test      [no test files]
```



> 也可以使用Golang的TDD小工具：`goconvey`
>
> 安装：`go get github.com/smartystreets/goconvey`
>
> 介绍: [GoConvey is awesome Go testing](https://github.com/smartystreets/goconvey)
>
> 运行：`goconvey`
>
> 效果截图：
>
> ![image-20181222173117732](/images/image-20181222173117732-5471077.png)



**测试用例**

创建`main_test.go`，

```bash
touch main_test.go
```

编写第一个测试用例：

```go
func TestHelloHandleFunc(t *testing.T) {
	rw := httptest.NewRecorder()
	name := "zouying"
	req := httptest.NewRequest(http.MethodPost, "/hello?name="+name, nil)
	handleHello(rw, req)

	if rw.Code != http.StatusOK {
		t.Errorf("status code not ok, status code is %v", rw.Code)
	}

	if len(counter) != 1 {
		t.Errorf("counter len not correct")
	}

	if counter[name] != 1 {
		t.Errorf("counter value is error: visitor=%s count=%v", name, counter[name])
	}
}
```



运行测试：`go test -v`：

> ➜  how_to_test git:(how_to_test) ✗ go test -v
> === RUN   TestHelloHandleFunc
> INFO[0000] visited                                       count=1 module=main name=zouying
> --- PASS: TestHelloHandleFunc (0.00s)
> PASS
> ok      _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test      0.015s



运行测试覆盖：`go test -cover`

> ➜  how_to_test git:(how_to_test) ✗ go test -cover
> INFO[0000] visited                                       count=1 module=main name=zouying
> PASS
> coverage: 62.5% of statements
> ok      _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test      0.021s



查看覆盖的代码：

```bash
#!/bin/bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```



> ➜  how_to_test git:(how_to_test) ✗ go test -coverprofile=/tmp/coverage.out
> INFO[0000] visited                                       count=1 module=main name=zouying
> PASS
> coverage: 62.5% of statements
> ok      _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test      0.015s
> ➜  how_to_test git:(how_to_test) ✗ go tool cover -html=/tmp/coverage.out

效果图为：

![image-20181222180733266](/static/images/image-20181222180733266-5473253.png)



绿色的表示测试代码覆盖住的，红色的表示没有覆盖。



第一个测试用例是直接测试http处理函数，我们使用了`httptest.NewRecorder()`创建`ResponseRecorder`对象，其中实现了 `ResponseWriter interface`。该对象在内存中记录了http response的状态。



还有一种测试方法是运行一个HTTP Server，使用HTTP Client请求该Server对应的接口。

httptest package中提供了`NewServer`方法，监听HandlerFunc处理函数，启动Server，启动Server的地址通过`URL`成员获得，例如：`http://127.0.0.1:52412`。需要注意的是，使用完毕后记得调用关闭：`Close()`。

代码如下，

```go
func TestHTTPServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handleHello))
	defer ts.Close()

	logrus.Infof("server url: %s", ts.URL)

	testURL := ts.URL + "/hello?name=zouying"
	resp, err := http.Get(testURL)
	if err != nil {
		t.Error(err)
		return
	}
	if g, w := resp.StatusCode, http.StatusOK; g != w {
		t.Errorf("status code = %q; want %q", g, w)
		return
	}
}
```

运行测试，

```bash
➜  how_to_test git:(master) ✗ go test -v -run=TestHTTPServer
=== RUN   TestHTTPServer
INFO[0000] server url: http://127.0.0.1:52506
INFO[0000] visited                                       count=1 module=main name=zouying
--- PASS: TestHTTPServer (0.00s)
PASS
ok      _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test      0.015s
```



**测试技巧：表格测试 (Table Based Tests)**



代码如下，

```go
func TestHelloHandlerMultiple(t *testing.T) {
	tests := []struct {
		name string
		wCnt int
	}{
		{name: "zouying", wCnt: 1},
		{name: "zouying", wCnt: 2},
		{name: "user2", wCnt: 1},
		{name: "user3", wCnt: 1},
	}

	for _, tc := range tests {
		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/hello?name="+tc.name, nil)
		handleHello(rw, req)

		if rw.Code != http.StatusOK {
			t.Errorf("status code not ok, status code is %v", rw.Code)
		}

		if counter[tc.name] != tc.wCnt {
			t.Errorf("counter value is error: visitor=%s count=%v", tc.name, counter[tc.name])
		}
	}
}
```



运行测试，

```bash
➜  how_to_test git:(how_to_test) ✗ go test -run=TestHelloHandlerMultiple
INFO[0000] visited                                       count=1 module=main name=zouying
INFO[0000] visited                                       count=2 module=main name=zouying
INFO[0000] visited                                       count=1 module=main name=user2
INFO[0000] visited                                       count=1 module=main name=user3
PASS
ok      _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test      0.016s
```



**测试工具：[testify](github.com/stretchr/testify/assert)**

使用工具介绍各种 `if {}`判断，产生大量的冗余代码。



代码，

```go
func TestHelloHandlerMultipleWithAssert(t *testing.T) {

	tests := []struct {
		name string
		wCnt int
	}{
		{name: "zouying", wCnt: 1},
		{name: "zouying", wCnt: 2},
		{name: "user2", wCnt: 1},
		{name: "user3", wCnt: 1},
	}

	for _, tc := range tests {
		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/hello?name="+tc.name, nil)
		handleHello(rw, req)

		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Equal(t, tc.wCnt, counter[tc.name])
	}
}
```



**Sub Test**



```go
func TestHelloHandlerInSubtest(t *testing.T) {

	tests := []struct {
		name string
		wCnt int
	}{
		{name: "zouying", wCnt: 1},
		{name: "user2", wCnt: 1},
		{name: "user3", wCnt: 1},
	}

	for _, tc := range tests {
		t.Run("test-"+tc.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/hello?name="+tc.name, nil)
			handleHello(rw, req)

			assert.Equal(t, http.StatusOK, rw.Code)
			assert.Equal(t, tc.wCnt, counter[tc.name])
		})
	}
}
```



运行测试，

```go
➜  how_to_test git:(how_to_test) ✗ go test -v . -run=TestHelloHandlerInSubtest
=== RUN   TestHelloHandlerInSubtest
=== RUN   TestHelloHandlerInSubtest/test-zouying
time="2018-12-23T23:07:19+08:00" level=info msg=visited count=1 module=main name=zouying
=== RUN   TestHelloHandlerInSubtest/test-user2
time="2018-12-23T23:07:19+08:00" level=info msg=visited count=1 module=main name=user2
=== RUN   TestHelloHandlerInSubtest/test-user3
time="2018-12-23T23:07:19+08:00" level=info msg=visited count=1 module=main name=user3
--- PASS: TestHelloHandlerInSubtest (0.00s)
    --- PASS: TestHelloHandlerInSubtest/test-zouying (0.00s)
    --- PASS: TestHelloHandlerInSubtest/test-user2 (0.00s)
    --- PASS: TestHelloHandlerInSubtest/test-user3 (0.00s)
PASS
ok      _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test      0.016s
```



### Data Race Detect

多个goroutine同时访问共享数据时，如果数据不是线程安全的，那么有可能会产生data race。

**HOW TO**

```bash
go test -race
```



**测试代码**

```bash
func TestHelloHandlerDetectDataRace(t *testing.T) {

	tests := []struct {
		name string
		wCnt int
	}{
		{name: "zouying", wCnt: 1},
		{name: "zouying", wCnt: 2},
		{name: "user2", wCnt: 1},
		{name: "user3", wCnt: 1},
	}

	for _, tc := range tests {
		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/hello?name="+tc.name, nil)
		handleHello(rw, req)

		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Equal(t, tc.wCnt, counter[tc.name])
	}
}
```



**运行测试**

```bash
➜  how_to_test git:(how_to_test) ✗ go test -race -v . -run=TestHelloHandlerDetectDataRace
=== RUN   TestHelloHandlerDetectDataRace
time="2018-12-23T22:58:22+08:00" level=info msg=visited count=1 module=main name=zouying
time="2018-12-23T22:58:22+08:00" level=info msg=visited count=2 module=main name=zouying
time="2018-12-23T22:58:22+08:00" level=info msg=visited count=1 module=main name=user2
time="2018-12-23T22:58:22+08:00" level=info msg=visited count=1 module=main name=user3
--- PASS: TestHelloHandlerDetectDataRace (0.00s)
PASS
ok      _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test      1.029s
```



测试通过，是否证明了我们的代码是没有问题的呢？



其实并非如此，只是没有检测出来。为什么没有检测出来？



是因为没有多个goroutine同时运行，访问共同的数据。



**修改代码**

```go
func TestHelloHandlerDetectDataRace(t *testing.T) {

	tests := []struct {
		name string
		wCnt int
	}{
		{name: "zouying", wCnt: 1},
		{name: "user2", wCnt: 1},
		{name: "user3", wCnt: 1},
	}

	var wg sync.WaitGroup
	wg.Add(len(tests))
	for _, tc := range tests {
		name := tc.name
		want := tc.wCnt

		go func() {
			defer wg.Done()

			rw := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/hello?name="+name, nil)
			handleHello(rw, req)

			assert.Equal(t, http.StatusOK, rw.Code)
			assert.Equal(t, want, counter[name])
		}()
	}
	wg.Wait()
}
```



**运行测试**

```bash
➜  how_to_test git:(how_to_test) ✗ go test -race . -run=TestHelloHandlerDetectDataRace
==================
WARNING: DATA RACE
Write at 0x00c0000a8f90 by goroutine 8:
  runtime.mapassign_faststr()
      /usr/local/go/src/runtime/map_faststr.go:190 +0x0
  _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test.handleHello()
      /Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test/main.go:14 +0x11c
  _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test.TestHelloHandlerDetectDataRace.func1()
      /Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test/main_test.go:144 +0x211

Previous read at 0x00c0000a8f90 by goroutine 7:
  runtime.mapaccess1_faststr()
      /usr/local/go/src/runtime/map_faststr.go:12 +0x0
  _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test.handleHello()
      /Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test/main.go:14 +0xbc
  _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test.TestHelloHandlerDetectDataRace.func1()
      /Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test/main_test.go:144 +0x211

Goroutine 8 (running) created at:
  _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test.TestHelloHandlerDetectDataRace()
      /Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test/main_test.go:139 +0x154
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:827 +0x162

Goroutine 7 (running) created at:
  _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test.TestHelloHandlerDetectDataRace()
      /Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test/main_test.go:139 +0x154
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:827 +0x162
==================
==================
WARNING: DATA RACE
Read at 0x00c0000a8f90 by goroutine 9:
  runtime.mapaccess1_faststr()
      /usr/local/go/src/runtime/map_faststr.go:12 +0x0
  _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test.handleHello()
      /Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test/main.go:14 +0xbc
  _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test.TestHelloHandlerDetectDataRace.func1()
      /Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test/main_test.go:144 +0x211

Previous write at 0x00c0000a8f90 by goroutine 8:
  runtime.mapassign_faststr()
      /usr/local/go/src/runtime/map_faststr.go:190 +0x0
  _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test.handleHello()
      /Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test/main.go:14 +0x11c
  _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test.TestHelloHandlerDetectDataRace.func1()
      /Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test/main_test.go:144 +0x211

Goroutine 9 (running) created at:
  _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test.TestHelloHandlerDetectDataRace()
      /Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test/main_test.go:139 +0x154
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:827 +0x162

Goroutine 8 (running) created at:
  _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test.TestHelloHandlerDetectDataRace()
      /Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test/main_test.go:139 +0x154
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:827 +0x162
==================
==================
WARNING: DATA RACE
Write at 0x00c0000a8f90 by goroutine 7:
  runtime.mapassign_faststr()
      /usr/local/go/src/runtime/map_faststr.go:190 +0x0
  _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test.handleHello()
      /Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test/main.go:14 +0x11c
  _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test.TestHelloHandlerDetectDataRace.func1()
      /Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test/main_test.go:144 +0x211

Previous write at 0x00c0000a8f90 by goroutine 8:
  runtime.mapassign_faststr()
      /usr/local/go/src/runtime/map_faststr.go:190 +0x0
  _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test.handleHello()
      /Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test/main.go:14 +0x11c
  _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test.TestHelloHandlerDetectDataRace.func1()
      /Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test/main_test.go:144 +0x211

Goroutine 7 (running) created at:
  _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test.TestHelloHandlerDetectDataRace()
      /Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test/main_test.go:139 +0x154
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:827 +0x162

Goroutine 8 (running) created at:
  _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test.TestHelloHandlerDetectDataRace()
      /Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test/main_test.go:139 +0x154
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:827 +0x162
==================
time="2018-12-23T23:13:06+08:00" level=info msg=visited count=1 module=main name=user2
time="2018-12-23T23:13:06+08:00" level=info msg=visited count=1 module=main name=user3
time="2018-12-23T23:13:06+08:00" level=info msg=visited count=1 module=main name=zouying
--- FAIL: TestHelloHandlerDetectDataRace (0.00s)
    testing.go:771: race detected during execution of test
FAIL
FAIL    _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test      0.030s
```



**分析报错**

```bash
==================
WARNING: DATA RACE
Write at 0x00c0000a8f90 by goroutine 8:
  runtime.mapassign_faststr()
      /usr/local/go/src/runtime/map_faststr.go:190 +0x0
  _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test.handleHello()
      /Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test/main.go:14 +0x11c
  _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test.TestHelloHandlerDetectDataRace.func1()
      /Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test/main_test.go:144 +0x211

Previous read at 0x00c0000a8f90 by goroutine 7:
  runtime.mapaccess1_faststr()
      /usr/local/go/src/runtime/map_faststr.go:12 +0x0
  _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test.handleHello()
      /Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test/main.go:14 +0xbc
  _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test.TestHelloHandlerDetectDataRace.func1()
      /Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test/main_test.go:144 +0x211
```

- goroutine 8, goroutine 7, ...
- DATA RACE
- `how_to_test/main.go:14`：`counter[name]++`
- `runtime/map_faststr.go:190`



原因是因为在多个goroutine中，对map同时进行了++操作，而在go中，map又不是线程安全的（线程安全的map参考sync包中的map），需要进行保护。



**修复race**



如果咱们的代码中有data race，那么一般使用下面方式可以避免，

- 使用channel。

  - [Share Memory By Communicating](https://blog.golang.org/share-memory-by-communicating)

  - [Go by Example: Channels](https://gobyexample.com/channels)

    - > ```go
      > messages := make(chan string, 2)
      > 
      > messages <- "buffered"
      > messages <- "channel"
      > 
      > ```

- 使用mutex。

  - [Package sync](https://golang.org/pkg/sync/)
  - [sync.Mutex - Tour of Go](https://tour.golang.org/concurrency/9)
  - [Go by Example: Mutexes](https://gobyexample.com/mutexes)

- 使用atomic。

  - [Go by Example: Atomic Counters](https://gobyexample.com/atomic-counters)

    > ```go
    > import "sync/atomic"
    > 
    > var ops uint64
    > 
    > for i := 0; i < 50; i++ {
    >     go func() {
    >         for {
    >             atomic.AddUint64(&ops, 1)
    >         }
    >     }()
    > }
    > 
    > opsFinal := atomic.LoadUint64(&ops)
    > ```

  - [sync/atomic package](https://golang.org/pkg/sync/atomic/)



**引入mutex解决问题**

1. 增加`var mu sync.Mutex`对`counter map`进行保护。
2. 在对counter访问前进行`Lock`操作，访问结束后，进行`Unlock`操作。



修改代码为，

```go
package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
)

var counter = map[string]int{}
var mu sync.Mutex // mutex for counter

func handleHello(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	mu.Lock()
	counter[name]++
	cnt := counter[name]
	mu.Unlock()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<h1 style='color: " + r.FormValue("color") +
		"'>Welcome!</h1> <p>Name: " + name + "</p> <p>Count: " + fmt.Sprint(cnt) + "</p>"))

	logrus.WithFields(logrus.Fields{
		"module": "main",
		"name":   name,
		"count":  cnt,
	}).Infof("visited")
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	http.HandleFunc("/hello", handleHello)
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}
```



测试代码，

```bash
➜  how_to_test git:(how_to_test) ✗ go test -v -race . -run=TestHelloHandlerDetectDataRace
=== RUN   TestHelloHandlerDetectDataRace
time="2018-12-24T10:11:23+08:00" level=info msg=visited count=1 module=main name=user3
time="2018-12-24T10:11:23+08:00" level=info msg=visited count=1 module=main name=zouying
time="2018-12-24T10:11:23+08:00" level=info msg=visited count=1 module=main name=user2
--- PASS: TestHelloHandlerDetectDataRace (0.00s)
PASS
ok      _/Users/zouying/src/Github.com/ZOUYING/learning_golang/how_to_test      1.025s
```



具体参考：

- https://golang.org/pkg/testing/