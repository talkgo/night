---
title: 2019-07-17 go程序运行初始化顺序
date: 2019-07-17T00:00:00+08:00
---
来源：『Go夜读』微信群

时间：2019-07-17

---

## 1. go程序运行初始化顺序

go程序运行初始化顺序如图：

![](/static/images/2019-07-17-go-program-init-sequence.jpg)



其实以前在看书时就看到过，不过当时并没有什么体会，通过下面这个问题代码，可以让读者对这个初始化顺序更有体会。

## 2. 问题代码描述

工程结构如图：

![](/static/images/2019-07-17-code01.png)

config/config.go:

```go
package config

import (
   "encoding/json"
   "io/ioutil"
   "log"
)

var (
   // 单例
   E_config *Config
)

type Config struct {

   MongodbUri string `json:"mongodbUri"`
   MongodbConnectTimeout int `json:"mongodbConnectTimeout"`
   LogBatchSize int `json:"logBatchSize"`
   LogCommitTimeout int `json:"logCommitTimeout"`

   NodeAddress string `json:"nodeAddress"`
   GenesisBlockMsg string `json:"genesisBlockMsg"`
   BlockChainDatabaseEngine string `json:"blockChainDatabaseEngine"`
   BlockChainDatabasePathTemplate string `json:"blockChainDatabasePathTemplate"`
}

// 读取配置
func InitConfig(configFile string) (err error) {

   var (
      content []byte
      config Config
   )

   // 读取配置文件，得到[]byte内容
   if content, err = ioutil.ReadFile(configFile); err != nil {
      log.Println("读取配置失败")
      return
   }

   // 反序列化
   if err = json.Unmarshal(content, &config); err != nil {
      return
   }

   // 赋值单例
   E_config = &config

   //log.Print(E_config)

   return

}
```

ledger/blockchain/chain.go:

```go
package blockchain

import (
	"log"
)

func InitChain() {

	log.Println(dbEngine, dbPathTemp, nodeAddress)

}
```

ledger/blockchain/config.go:

```go
package blockchain

import "../../config"

var (
   // json配置
   dbPathTemp = config.E_config.BlockChainDatabasePathTemplate
   dbEngine = config.E_config.BlockChainDatabaseEngine
   nodeAddress = config.E_config.NodeAddress
)
```

ledger/cli/cli.go:

```go
package main

import (
   "../../config"
   "../blockchain"
   "log"
)

func initConfig() {
   const configFile = "config/enode.json"

   if err := config.InitConfig(configFile); err != nil {
      log.Panic(err)
   }
}

func main() {
   initConfig()
   blockchain.InitChain()
}
```

我的最初想法是在blockchain下新建config.go，将config.E_config的元素赋给config.go/dbEngine等，将这些从外部包引用的配置参数集中放置，达到提升代码可读性的目的，但是这似乎不行：

```
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x0 addr=0x58 pc=0x4cd34d]

goroutine 1 [running]:
_/E_/GO/XProjects/EChain-error/ledger/blockchain.init.ializers()
	E:/GO/XProjects/EChain-error/ledger/blockchain/config.go:13 +0x2d

Process finished with exit code 2
```

程序直接panic掉了，报错指向ledger/blockchain/config.go下dbEngine等变量的初始化，并指出产生了空指针调用。

但是看上去好像没有问题啊？main程序先执行配置初始化，得到E_config，再初始化Chain，Chain中使用E_config得到的元素值。

问题出在了变量初始化的顺序。根据go程序初始化顺序，一个程序在执行时，首先会去寻找import，然后会进行全局常量、全局变量然后是init()函数的初始化，最后才会执行到main()函数（也就是常说的程序入口）。

在我的代码中，dbEngine等变量就是blockchain中的包内全局变量，因此程序应该是先进行dbEngine等的初始化赋值，然而，此时尚未执行到main()，也就不会执行initConfig()，E_config也就仍然不存在。因此在dbEngine等元素赋值时我赋了一个根本不存在的值给到它们，这显然是无效的内存地址。

知道问题出在哪以后，就可以进行改正了，最简单的改正措施就是，将ledger/blockchain/config.go删掉，将dbEngine等变量初始化赋值语句放到InitChain中，使之由全局变量转为局部变量，这样，其初始化赋值操作就在InitConfig()之后了，也就不会产生空指针调用的错误。更正后代码及运行结果如下：

ledger/blockchain/chain.go:

```go
package blockchain

import (
   "../../config"
   "log"
)


func InitChain() {

   var (
      // json配置
      dbPathTemp = config.E_config.BlockChainDatabasePathTemplate
      dbEngine = config.E_config.BlockChainDatabaseEngine
      nodeAddress = config.E_config.NodeAddress
   )
   
   log.Println(dbEngine, dbPathTemp, nodeAddress)

}
```

运行ledger/cli/cli.go/main()的结果:

```
2019/07/18 11:20:13 badger ./tmp/blocks/%s/blocks_%s 127.0.0.1:9797

Process finished with exit code 0
```

这确实是从enode.json中读取的配置信息。说明现在没有问题了。

## 3. 起因及鸣谢

作为一个Go语言初学者及区块链初学者，在掌握了一些前置基础之后，我开始了EChain项目(一个基于区块链进行设备管理的物联网平台？)的开发（其实也是导师布置的任务）。在开发的过程中遇到了本文中提到的问题，非常感谢**夜读群中Ryan等人**热心的指导与帮助，在此谢过！另外希望EChain能够顺利完成！（来自菜鸟的祈祷~）

