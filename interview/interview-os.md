# 操作系统

_1.Select，Poll，Epoll的区别？_

> `select`，`poll`，`epoll`都是IO多路复用的机制，具体区别请查阅资料   
>
> [查看资料](https://blog.csdn.net/windeal3203/article/details/52055436)

_2.什么叫虚拟内存？_

> 虚拟内存是计算机系统内存管理的一种技术。它使得应用程序认为它拥有连续的
可用的内存（一个连续完整的地址空间），而实际上，它通常是被分隔成多个物理内
存碎片，还有部分暂时存储在外部磁盘存储器上，在需要时进行数据交换。   

_3.什么叫桥接？_

> 桥接是指依据OSI网络模型的链路层的地址，对网络数据包进行转发的过程，工作
在OSI的第二层；一般的交换机，网桥就有桥接作用。        

_4.Linux什么命令可以查看cpu和内存？怎么查看每个核的cpu呢？_

> top命令        
> 在top查看界面按数字1即可查看每个核的数据        

_5.给一个PID=100你觉得它是后台程序还是前台程序？_

> 进程号0-299保留给daemon进程         

_6.怎么查看一个端口的TCP连接情况？_

> netstat          

_7.Docker的网络模式有哪几种？_

> bridge网络          
> host网络         
> none网络         
> container模式        

_8.介绍一下Tcpdump？_

> tcpdump网络数据包截获分析工具。支持针对网络层、协议、主机、网络或端口
的过滤。并提供and、or、not等逻辑语句帮助去除无用的信息。      
>
> [查看资料](https://www.cnblogs.com/chyingp/p/linux-command-tcpdump.html)

_9. 什么叫大端和小端？_

> **说明**
>> 1.Little-Endian（小端）就是低位字节排放在内存的低地址端，高位字节排放在内存的高地址端        
>> 2.Big-Endian（大端）就是高位字节排放在内存的低地址端，低位字节排放在内存的高地址端。
>
> **使用场景**
>
>> 一般`操作系统`都是小端的，而`通信协议`是大端的         
>
>> [查看资料](https://blog.csdn.net/element137/article/details/69091487)

_10. 介绍下 docker 底层原理_

> 1. [查看资料](https://draveness.me/docker)
> 2. [左耳朵耗子 Docker 基础技术介绍(有例程)](https://coolshell.cn/tag/docker)

_11.介绍些僵尸进程和孤儿进程的区别, 怎么产生的, 怎么避免?_

> [查看资料](https://www.cnblogs.com/lxmhhy/p/6212405.html)

_12.CPU 使用率和 CPU 负载有什么区别?_
> [查看资料](https://www.cnblogs.com/muahao/p/6492665.html)