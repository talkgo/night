有下面一段程序，面试官问这段程序有什么问题？

```
type Store struct {
    a string
    b string

    sync.RWMutex
}

func (s *Store) GetA() string {
    fmt.Println("get a")
    s.RLock()
    fmt.Println("get a2")
    defer s.RUnlock()

    return s.a
}

func (s *Store) GetAB() (string, string) {
    fmt.Println("get ab")
    s.RLock()
    fmt.Println("get ab2")
    defer s.RUnlock()

    return s.GetA(), s.b
}

func (s *Store) Write(a, b string) {
    fmt.Println("write")
    s.Lock()
    defer s.Unlock()
    fmt.Println("write2")

    s.a = a
    s.b = b
}
```

1. 看到这段程序程序，首先想到的是读写锁的问题；
2. 其次，看 Store 这个结构体，各个函数都定义的是指针函数。那就说明：不存在读写锁的 copy 过程；
3. GetAB 方法中通过调用 GetA 方法，在 `s.RUnlock` 前通过调用 `s.GetA`，又做了一次读写锁上锁 `s.RLock`，但是读锁可以多次上锁，所以单看这里没什么问题；
4. 然后，想到会不会 `Write` 和 `GetAB` 并发调用的时候会存在问题呢？思考了一会，觉得没问题，就放弃了。

以上，是面试时整个思路。

回头，越想越觉得这里哪里有问题，就在`夜读群`里求教了一下，群里大神发了一篇[读写锁优先级的文章](https://blog.csdn.net/xyz347/article/details/83902123)，然后给了一段测试样例，瞬间豁然开朗。

main函数逻辑如下：

```
func main() {
    store := Store{}
    wg := sync.WaitGroup{}
    wg.Add(2)
    go func() {
        defer wg.Done()
        for i := 1; i < 10000; i += 1 {
            fmt.Println("main write ", i)
            store.Write("111", "1111")
        }
    }()
    go func() {
        defer wg.Done()
        for i := 1; i < 10000; i += 1 {
            fmt.Println("main get ab", i)
            store.GetAB()
        }
    }()
    wg.Wait()
}
```

执行结果

```
main get ab 12  //main函数读取ab
get ab          //进入 s.GetAB 函数
main write  1   //main 函数写数据
write           //进入 s.Write 函数
write2          //获取写锁
main write  2
....            //写锁一直抢占
....
main write  13  //main 函数写数据
write           //进入 s.Write 函数
write2          //获取写锁
main write  14  
write
write2
get ab2         //之前 get ab 12 才获得读锁
get a           //进入 GetA 
get a2          //获取读锁
main get ab 13  //main函数 get ab
get ab          //进入 s.GetAB 函数
get ab2         //获取读锁
main write  15  //注意⚠️ 这个时候写数据开始了
write           //进入 Write 函数，后面尝试获取写锁
get a           //这个时候 GetAB 进入了 GetA，尝试获取读锁
fatal error: all goroutines are asleep - deadlock!  //出现了死锁
```

分析：

```

    GetAB       |            GetA               |                Write
                |                               |
r0 占用读锁      |                               |
                |                               |    w0 尝试获取写锁 等待r0释放读锁
                |  r1 尝试获取读锁，排在w0后面      |

```

由于读写锁的优先级，读锁和写锁同时竞争时，读锁要排在写锁后面，导致了 r1 竞争 w0的锁，w0竞争r0，r0执行不下去，最后死锁。

所以，在以后用锁的时候不管有没有优先级，都要时刻记住死锁的四个必要条件：

1. 互斥条件：一个资源每次只能被一个进程使用。 
2. 锁的不可抢占：进程已获得的资源，在末使用完之前，不能强行剥夺。
3. 占有且等待：一个进程因请求资源而阻塞时，对已获得的资源保持不放。 
4. 循环等待条件: 若干进程之间形成一种头尾相接的循环等待资源关系。

