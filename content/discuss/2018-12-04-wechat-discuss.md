---
title: "2018-12-04 微信讨论"
date: 2018-12-04T17:13:46+08:00
---

来源：『Go 夜读』微信群

## json 解析结构体断言问题

```golang
package main

import (
	"encoding/json"
)

type Msg struct {
	Id     int           `json:"id"`
	Params []interface{} `json:"params"`
}

type Employee struct {
	Name string `json:"name"`
}

func main() {
	b := []byte(`{"id":111,"params":["boss",[{"name":"a"},{"name":"b"}]]}`)
	msg := &Msg{}
	json.Unmarshal(b, msg)

	_ = msg.Params[0].(string)
	_ = msg.Params[1].([]Employee) //无法直接断言?
}

```

对于复杂多变的结构体, 经常会用 interface 作为类型,然后根据情况断言

上面的例子运行起来会报错, 原因是 `json.Unmarshal` 之后无法再断言成具体的结构体 (虽然结构体和 map[string]interface{}可以互转)

### 解决方案1️⃣(非最优):

把想要断言成具体 struct 的字段重新`json.Marshal` 和 `json.Unmarshal`,会多操作一步不推荐

```golang
newB, _ := json.Marshal(msg.Params[1])
var employee []Employee
json.Unmarshal(newB, &employee)
```

### 解决方案2️⃣:

把 interface{}类型改为 `json.RawMessage` ,可以延迟该字段的解析, 在用到的时候进行二次解析 

```golang
package main

import (
	"encoding/json"
	"log"
)

type Msg struct {
	Id     int               `json:"id"`
	Params []json.RawMessage `json:"params"`
}

type Employee struct {
	Name string `json:"name"`
}

func main() {
	b := []byte(`{"id":111,"params":["boss",[{"name":"a"},{"name":"b"}]]}`)
	msg := &Msg{}
	json.Unmarshal(b, msg)

	var boss string
	json.Unmarshal(msg.Params[0], &boss)
	log.Println(string(boss))

	var employee []Employee
	json.Unmarshal(msg.Params[1], &employee)
	log.Println(employee)
}

```

