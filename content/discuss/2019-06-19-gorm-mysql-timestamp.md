---
title: 2019-06-19 Go、Gorm 与 MySQL timestamp
date: 2019-06-19T00:00:00+08:00
---

来源：『Go 夜读』微信群

----

Go、Gorm与MySQL中timestamp交互时遇到的问题

涉及到的方面

- MySQL中timestamp 默认值，explicit_defaults_for_timestamp属性设置
- Go中time.Time字段类型，time.Time零值
- Gorm中的处理方式


例如:

数据模型：

```
type A struct {
	Id            int
	UserId        int
	VipExpireTime time.Time
	MessageType   string
	ClickTabTime  time.Time
	CreateTime time.Time `gorm:"default:current_time"`
	UpdateTime time.Time `gorm:"default:current_time"`
}
```

对应字段：

```
CREATE TABLE `a` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `vip_expire_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT 'vip终止时间',
  `message_type` varchar(50) NOT NULL DEFAULT '' COMMENT '消息的类型',
  `click_tab_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '点击Tab的时间',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
COMMENT='a表';
```

数据库初始化：

```
import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBOrm struct {
	Orm *gorm.DB
}

var DB DBOrm

const (
	dbTestHost = "127.0.0.1"
	dbTestUser = "root"
	dbTestPwd  = "123"
	dbDevDB    = "test1"
)

func InitGorm(user, password, addr, db string) {
	var err error
	DB.Orm, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, addr, db))
	if err != nil {
		panic(err)
	}

	DB.Orm.LogMode(true)
}

func InitDebug() {
	InitGorm(dbTestUser, dbTestPwd, dbTestHost, dbDevDB)
}
```

```
func timenull() {

	InitDebug()
	a := A{
		UserId:        1,
		// VipExpireTime: time.Time{},
	}
	err := DB.Orm.Table("a").Create(&a).Error
	fmt.Println(err)
	fmt.Printf("a: %+v", a)
}
```

执行结果：
```
d:\mygo\src\ch\t>go test -v -run TestTimenull
=== RUN   TestTimenull

?[35m(D:/mygo/src/ch/t/run.go:263)?[0m
?[33m[2019-06-18 16:41:16]?[0m  ?[36;1m[1.01ms]?[0m  INSERT  INTO `a` (`user_id`,`vip_expire_time`,`message_type`,`click_tab_time`) VALUES (1,'0001-01-01 00:00:00','','0001-01-01 00:00:00')
?[36;31m[1 rows affected or returned ]?[0m
<nil>
a: {Id:3 UserId:1 VipExpireTime:0001-01-01 00:00:00 +0000 UTC MessageType: ClickTabTime:0001-01-01 00:00:00 +0000 UTC}--- PASS: TestTimenull (0.02s)
PASS
ok      ch/t    0.266s
```

发现这样`gorm`中操作是可以创建成功的。但是，如果粘贴`insert`语句到数据库中执行，是报错的。

```
INSERT  INTO `a` (`user_id`,`vip_expire_time`,`message_type`,`click_tab_time`) VALUES (1,'0001-01-01 00:00:00','','0001-01-01 00:00:00')
```

```
错误代码： 1292
Incorrect datetime value: '0001-01-01 00:00:00' for column 'vip_expire_time' at row 1
```

造成这种时间的原因是什么呢？

大概是因为`Go`语言中`time`的初始值是`第一年的一月一日`这个设定。
[Golang Time](https://golang.org/pkg/time/#Time)

想避免这种方式要怎么处理呢？

也带着问题问在夜读群中讨论。

「杨文：@我的名字叫浩仔丶Go 请教一个关于gorm create的问题，结构体user内部有一个time.Time字段a，对应数据库是timestamp类型，create是如果没有对a赋值，insert会报错，插入时间为0001-01-01了。修改方案想了两种，一个是给a赋time.Time{}，另一种是将a改为指针的time，插入null。这两种哪种好呢，大家是怎么处理的呢？
@jinzhu gorm 作者」

「jinzhu：可以用 *time.Time ，或者类似 NullTime 这种类型」

「jinzhu：并且你的mysql应该是5.7之后的新版本吧，有个变量，允许 0001-01-01 这类数据。。。」

Gorm 作者提到的两种方式：

- `*time.Time`（这貌似也是gorm issue里大部分的答案）
- `NullTime`

第一种方式

```
type A struct {
	Id            int
	UserId        int
	VipExpireTime *time.Time
	MessageType   string
	ClickTabTime  *time.Time
}

a := A{
	UserId:        1,
}
	
INSERT  INTO `a` (`user_id`,`vip_expire_time`,`message_type`,`click_tab_time`) VALUES (1,NULL,'',NULL);
```

这样`gorm`中操作是报错的。

```
Error Code: 1048. Column 'vip_expire_time' cannot be null
```

另外此时又会引入新的问题，因为字段设置为指针类型，所以再取值时需要判断是否为null，否则会空指针。

在每一个用到`VipExpireTime`的地方，都需要判断

```
if VipExpireTime != nil { VipExpireTime.Format("2006-01-02 15:04:05") }
```

这种代码让人头大！

再有Go Time的定义，也不建议用*time.Time。

>Programs using times should typically store and pass them as values, not pointers. That is, time variables and struct fields should be of type time.Time, not *time.Time.

第二种方式：

```
// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return nil, nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC3339))
	return []byte(val), nil
}

type A struct {
	Id     int
	UserId int
	// VipExpireTime time.Time
	// VipExpireTime *time.Time
	VipExpireTime NullTime
	MessageType   string
	// ClickTabTime  time.Time
	// ClickTabTime *time.Time
	ClickTabTime NullTime
	CreateTime   time.Time `gorm:"default:current_time"`
	UpdateTime   time.Time `gorm:"default:current_time"`
}

type NullTime struct {
	mysql.NullTime
}
```

可以参照这篇文章做处理：[How I handled possible null values from database rows in Golang?](https://medium.com/aubergine-solutions/how-i-handled-null-possible-values-from-database-rows-in-golang-521fb0ee267)

然后我们说一下MySQL中[explicit_defaults_for_timestamp](https://dev.mysql.com/doc/refman/5.6/en/server-system-variables.html#sysvar_explicit_defaults_for_timestamp)属性，这与timestamp的默认值类型与表现形式有关。

>注意：explicit_defaults_for_timestamp本身已被弃用，因为它的唯一目的是允许控制将来在MySQL版本中删除的已弃用的TIMESTAMP行为。当删除这些行为时，explicit_defaults_for_timestamp将没有任何用途，也将被删除。

查看`explicit_defaults_for_timestamp`当前的状态：

```
SHOW VARIABLES LIKE 'explicit_defaults_for_timestamp';

explicit_defaults_for_timestamp: OFF
```

然后参考MySQL文档中给出的方式处理，[Automatic Initialization and Updating for TIMESTAMP and DATETIME](https://dev.mysql.com/doc/refman/5.6/en/timestamp-initialization.html)

因为`timestamp`会有 `create_time`、`update_time`这种字段，如果不赋值，gorm会按照零值处理，所以可以在字段后加`tag`

```
type A struct {
	Id            int
	UserId        int
	VipExpireTime time.Time
	MessageType   string
	ClickTabTime  time.Time
	CreateTime    time.Time `gorm:"default:current_time"`
	UpdateTime    time.Time `gorm:"default:current_time on update current_time"`
}

对应数据库字段：

CREATE TABLE `a` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `vip_expire_time` timestamp DEFAULT 0 COMMENT 'vip终止时间',
  `message_type` varchar(50) NOT NULL DEFAULT '' COMMENT '消息的类型',
  `click_tab_time` timestamp DEFAULT 0 COMMENT '点击Tab的时间',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
COMMENT='a表';
```

这样对于后续业务的时间判断，就可以利用 time.IsZero() 来判断。

看这个讨论[Should I use the datetime or timestamp data type in MySQL?](https://stackoverflow.com/questions/409286/should-i-use-the-datetime-or-timestamp-data-type-in-mysql)

当然直接把MySQL字段设置为datetime，也可以规避此类问题。但是对于使用datetime，还是timestamp？
个人还是倾向timestamp吧，因为随时区变化，空间效率更高。

扩展：

一、datetime 与 timestamp

***datetime***

范围最大，`1001`到`9999`年，时间格式为`YYYYMMDDHHMMSS`，与时区无关，使用**8个字节**存储。

如果没有指定 `default`，datetime 默认为 `null`。
```
Go中time.Time{}为零时，值为：0001-01-01 00:00:00 +0000 UTC
```

***timestamp***

保存了从`1970年1月1日午夜`以来的秒数，与UNIX时间戳相同。范围从1970年到2038年；使用**4个字节**存储；显示的值依赖于时区。

timestamp可以配置插入更新的行为，
如果没有指定 `default`，timestamp 默认为 `0`(即 1970-01-01 00:00:00)。

如果强行更新小于1970年的值，会报错：
```
Incorrect datetime value: '1969-12-01 00:00:00' for column 'ts' at row 1
```

二、UTC/GMT/时间戳

1.UTC时间 与 GMT时间

我们可以认为格林威治时间就是时间协调时间（GMT=UTC），格林威治时间和UTC时间均用秒数来计算的。

2.UTC时间 与 本地时

UTC + 时区差 ＝ 本地时间
时区差东为正，西为负。在此，把东八区时区差记为 +0800，

UTC + (＋0800) = 本地（北京）时间 (1)

那么，UTC = 本地时间（北京时间)）- 0800 (2)

3.UTC 与 Unix时间戳

在计算机中看到的UTC时间都是从（1970年01月01日 0:00:00)开始计算秒数的。所看到的UTC时间那就是从1970年这个时间点起到具体时间共有多少秒。 这个秒数就是Unix时间戳。


参考资料：

《高性能MySQL》

[Automatic Initialization and Updating for TIMESTAMP and DATETIME](https://dev.mysql.com/doc/refman/8.0/en/timestamp-initialization.html)

[How I handled possible null values from database rows in Golang?](https://medium.com/aubergine-solutions/how-i-handled-null-possible-values-from-database-rows-in-golang-521fb0ee267)