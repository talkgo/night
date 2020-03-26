---
desc: Go 夜读 && TiDB 源码阅读之概览
title: 第 46 期 TiDB 源码阅读之概览
date: 2019-06-05T21:10:00+08:00
author: 龙恒
---

## TiDB Source Code Overview

![image](https://user-images.githubusercontent.com/1710912/58966936-c7bfb100-87e5-11e9-9479-8f9e95105227.png)

## 视频回看

1. [TiDB 源码学习之 Source Code Overview - YouTube](https://youtu.be/mK6BOquvQhE)
2. [TiDB 源码学习之 Source Code Overview - Bilibili](https://www.bilibili.com/video/av54658699/)

## 意见反馈

1. [【Go夜读】『TiDB Source Code Overview』反馈](https://docs.google.com/forms/d/e/1FAIpQLSeaj0ZxZJhfqa0oS8MZGtTDIylSCAdLq1ymnkYhfbkgSQ6rOw/viewform)

## chat 答疑

20:54:52	 From mai yang : 大家好，欢迎大家前来参加 Go 夜读&TiDB 源码学习！
21:22:34	 From nange : Session 怎么初始化的？
21:22:46	 From ccong deng : 每个连接都是跟一个session对象对应么？
21:22:48	 From jeffery : session主要包含什么？
21:22:59	 From jeffery : 譬如：
21:23:01	 From Wei Yao : 对，一个链接一个 session
21:23:09	 From Wei Yao : 具体包含什么，可以大家自己去看了
21:23:17	 From Wei Yao : 这个线上不可能所有都讲的
21:23:25	 From jeffery : 好的，谢谢了
21:31:00	 From Wei Yao : 大家如果对语法分析，词法分析感兴趣，可以去看看 yacc 跟 lex
21:31:10	 From hezhiyong : parser 这一层不是使用mysql的parser吗 
21:31:17	 From Wei Yao : 不，我们自己写的
21:31:56	 From hezhiyong : mysql 的语法解析是在那一步用到了？
21:32:01	 From tianyi wang : select coalesce（）中coalesce是在fields里面吗
21:32:10	 From Wei Yao : 我们的语法解析就是兼容 mysql，
21:32:12	 From window930030@gmail.com : SQL injection 有做嗎？
21:32:27	 From Wei Yao : SQL injection？SQL 注入？
21:32:40	 From Wei Yao : 我们不叫 sql 注入
21:32:55	 From window930030@gmail.com : 恩？
21:33:04	 From Wei Yao : 我们会把 sql 变成算子，之后会去优化算子结构，下面会讲，
21:33:15	 From window930030@gmail.com : 好的，謝謝。
21:34:55	 From jeffery : 刚刚的意思：Visitor是选择节点
21:34:58	 From jeffery : ？
21:35:05	 From Wei Yao : 不是
21:35:11	 From Abner Zheng : 一种设计模式
21:35:13	 From xietengjin : 遍历节点用的吧
21:35:14	 From Wei Yao : visitor 是设计模式中的那个 visitor 模式
21:35:17	 From Wei Yao : 对
21:35:19	 From jacobz : 遍历树用的
21:35:23	 From Fangfang Qi : 是遍历语法树的
21:35:27	 From jeffery : 额，好的
21:35:28	 From Wei Yao : 遍历 ast 树
21:37:03	 From jeffery : 清楚
21:40:01	 From jacobz : 是搞优化的那一堆？
21:43:45	 From lk : 递归遍历？
21:44:05	 From Wei Yao : 层级有限。
21:48:06	 From Kathy : 其实这个时候是不是类似传统的通过运算符进栈出栈形成表达式
21:48:24	 From Wei Yao : 对，表达式系统基本上都是这样
21:51:21	 From Kathy : ScalarFunction能解决aggregation的函数的语句吗
21:55:24	 From Chen Shuang : 能
21:55:55	 From Chen Shuang : aggregation function 也是 scalar function.
21:57:23	 From Kathy : 只要不涉及其他表的相关列的function是否都最后成为scalarFunction
21:57:33	 From Kathy : 的表达式
21:58:40	 From Chen Shuang : 只要是 function , 都会变成 scalarFunction 表达式
21:59:40	 From Chen Shuang : select t1.a + t2.b from t1,t2; 其中 t1.a + t2.b 会build 成一个 scalarFunction 表达式
21:59:40	 From Kathy : 多谢答复
22:00:28	 From Chen Shuang : 不客气哈
22:06:26	 From tianyi wang : select coalesce（）也会是scalarfunction?
22:06:53	 From hezhiyong : 可以演示一下debug一条语句跑的代码吗
22:10:07	 From jiangchen : 是的，能不能最好演示下。。每次next返回的是一部分子结果还是一部分最终的结果？
22:11:33	 From Kathy : 执行引擎的新特性可以说说吗？简单讲一下，就是parallel physical operator的实现等等
22:11:57	 From Wei Yao : 执行引擎下周讲
22:12:05	 From 慢摇哥哥 : 老师，Coprocessor是在哪一步分发的
22:12:06	 From jeffery : 辛苦了，有一个基本的逻辑了
22:12:37	 From Kathy : 好的 谢谢
22:12:42	 From 达 黄 : 之前看了tidb源码解析的文章 配合着这个视频 印象更清楚了
22:13:32	 From jeffery : 感觉姚老师像一位老教授在督导
22:14:02	 From Wei Yao : ：）
22:15:40	 From jeffery : 为什么这部分会单独出来？
22:15:40	 From nange : Distsql是什么好像没讲。
22:15:54	 From tianyi wang : select coalesce（）会是scalarfunction还是单独的一部分呢?
22:18:53	 From 熊浪 : 问下是每一个session都会解析一次sql么？如果一个sql在同一个session中多次执行是否有ast的共享？
22:21:02	 From hezhiyong : prepare 是要开启参数才可以的吧
22:21:36	 From 熊浪 : 好的，和mysql是一样的。谢谢

PPT: https://talkgo.slack.com/files/U8A45L223/FKA335THT/_reading-go__tidb_source_cdoe_overview.pdf

## 观看视频

{{< youtube id="mK6BOquvQhE" >}}
