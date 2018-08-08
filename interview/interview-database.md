# 数据库

_1.Mysql事物的隔离级别?_

| **事务隔离级别**   |  **脏读** | **不可重复读** | **幻读** |
| ----- | ----- | ----- | ----- |
| 读未提交（read-uncommitted） | 是 | 是 | 是 |
| 读已提交（read-committed） | 否 | 是 | 是 |
| 可重复读（repeatable-read） | 否 | 否 | 是 |
| 串行化（serializable） | 否 | 否 | 否 |
>
> [相关资料](https://www.cnblogs.com/huanongying/p/7021555.html)

_2.Innodb和Myisam的区别？_

> Innodb支持事务，而Myisam不支持事务        
> Innodb支持行级锁，而Myisam支持表级锁        
> Innodb支持外键，而Myisam不支持       
> Innodb不支持全文索引，而Myisam支持          
> Innodb是索引组织表， Myisam是堆表
>
> [相关资料](https://blog.csdn.net/nuli888/article/details/52443011)         

_3.Mysql慢响应默认时间?_

> `10s`         

_4.Explain的含义?_

> explain显示了mysql如何使用索引来处理select语句以及连接表。可以帮助
选择更好的索引和写出更优化的查询语句。        

_5.Profile的意义以及使用场景?_

> Profile用来分析SQL性能的消耗分布情况。当用explain无法解决慢SQL的时
候，需要用profile来对SQL进行更细致的分析，找出SQL所花的时间大部分消耗在
哪个部分，确认SQL的性能瓶颈。        

_6.Redis的过期失效机制？_

> `scan`扫描+给每个key存储过期时间戳

_7.Redis持久化方案aof的默认fsync时间是多长？_

> `1s`     

_8.Redis持久化方案rdb和aof的区别？_

> [查看资料](https://juejin.im/post/5ab5f08e518825557f00dfac)

_9.Redis怎么查看延迟数据?（非业务操作）_

> **可以用redis-cli工具加--latency参数可以查看延迟时间**
>       
>> `redis-cli --latency -h 127.0.0.1 -p 6379`     
>  
> **使用slowlog查出引发延迟的慢命令**
> 
>> `slowlog get`       

![](./images/slowlog.jpeg)

_10.Redis的集群怎么搭建？_

> [查看资料](https://segmentfault.com/a/1190000008448919)

_11.简单介绍下什么是缓存击穿, 缓存穿透, 缓存雪崩? 能否介绍些应对办法?_

> [查看资料](https://blog.csdn.net/zeb_perfect/article/details/54135506)

_12.关系型数据库 MySQL/PostgreSQL的索引类型? 其他数据库优化方法?_

> 1. https://segmentfault.com/a/1190000003072424
> 2. https://tech.meituan.com/performance_tunning.html

_13.介绍下数据库分库分表以及读写分离?_

> 分库分表主要解决写的问题, 读写分离主要解决读的问题.
> 分库分表的策略有很多种: 平均分配, 按权重分配, 按业务分配, 一致性 hash.....  
> 读写分离的原理大致是一台主、多台从。主提供写操作，从提供读操作.   
> 方案可以根据以下几个因素来综合考虑:     
> > 1.数据实时性要求？  
> > 2.查询复杂度是否比较高？  
> > 3.读和写的比例即侧重点是哪一个？  
> 
> 方案有很多, 大家可以自行搜索, 学习总结. 
这题源自微博的平台技术专家一条微博: https://m.weibo.cn/status/4265027340366901