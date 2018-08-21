# 数据结构

_1.什么是跳跃表?_

> 跳跃表是基于有序链表的一种扩展        
>
> [查看资料](http://blog.jobbole.com/111731/)

_2. 介绍下 RESTFull API 方式下, 怎么做到快速路由?_

> 一般使用前缀树/字典树, 来提高查找速度. 

> 开源的路由模块里, httprouter 是比较快的, 参考: https://github.com/julienschmidt/httprouter   
> 他使用了一种改进版的前缀树算法. 这个树的应用非常广泛, 除了做路由, 还有 linux 内核里使用, 在数据库里也有用到.   
> 参考文章:   
> 1. [路由查找之Radix Tree](https://michaelyou.github.io/2018/02/10/%E8%B7%AF%E7%94%B1%E6%9F%A5%E6%89%BE%E4%B9%8BRadix-Tree/)  
> 2. [图文详解Radix树](https://blog.csdn.net/petershina/article/details/53313624)  
> 3.[radix tree在数据库PostgreSQL中的一些应用举例](https://yq.aliyun.com/articles/75334)