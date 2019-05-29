、（-.-）、

reflect.Value 类型有一个 UnsafeAddr 方法。这个方法会返回一个 uintptr 类型的结果值。

这种行为是————————————>不安全<————————————————————的！！！！！！

UnsafeAddr 方法返回的指针值所指向的内存地址很可能会在之后的某个时刻被存入其他的东西，从而使得你对内存中数据的直接更改导致整个程序的崩溃。

这是一个被拒绝的proposal  https://github.com/golang/go/issues/19752
