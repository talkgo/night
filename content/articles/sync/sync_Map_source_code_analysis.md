---
title: sync.Map源码分析
---
## 背景
众所周知,go普通的map是不支持并发的，换而言之,不是线程(goroutine)安全的。博主是从golang 1.4开始使用的，那时候map的并发读是没有支持，但是并发写会出现脏数据。golang 1.6之后，并发地读写会直接panic：
```
fatal error: concurrent map read and map write
```

```go
package main
func main() {
	m := make(map[int]int)
	go func() {
		for {
			_ = m[1]
		}
	}()
	go func() {
		for {
			m[2] = 2
		}
	}()
	select {}
}
```
所以需要支持对map的并发读写时候，博主使用两种方法：
1. 第三方类库 [concurrent-map](https://github.com/orcaman/concurrent-map)。
2. map加上sync.RWMutex来保障线程(goroutine)安全的。

golang 1.9之后,go 在sync包下引入了并发安全的map，也为博主提供了第三种方法。本文重点也在此，为了时效性，本文基于golang 1.10源码进行分析。

## sync.Map
### 结构体
#### Map
```go
type Map struct {
	mu Mutex    //互斥锁，用于锁定dirty map

	read atomic.Value //优先读map,支持原子操作，注释中有readOnly不是说read是只读，而是它的结构体。read实际上有写的操作

	dirty map[interface{}]*entry // dirty是一个当前最新的map，允许读写

	misses int // 主要记录read读取不到数据加锁读取read map以及dirty map的次数，当misses等于dirty的长度时，会将dirty复制到read
}
```

#### readOnly
readOnly 主要用于存储，通过原子操作存储在Map.read中元素。
```
type readOnly struct {
	m       map[interface{}]*entry
	amended bool // 如果数据在dirty中但没有在read中，该值为true,作为修改标识
}
```

#### entry

```
type entry struct {
	// nil: 表示为被删除，调用Delete()可以将read map中的元素置为nil
	// expunged: 也是表示被删除，但是该键只在read而没有在dirty中，这种情况出现在将read复制到dirty中，即复制的过程会先将nil标记为expunged，然后不将其复制到dirty
	//  其他: 表示存着真正的数据
	p unsafe.Pointer // *interface{}
}
```

### 原理
如果你接触过大Java，那你一定对CocurrentHashMap利用**锁分段技术**增加了锁的数目，从而使争夺同一把锁的线程的数目得到控制的原理记忆深刻。  
那么Golang的sync.Map是否也是使用了相同的原理呢？sync.Map的原理很简单，使用了**空间换时间**策略，通过冗余的两个数据结构(read、dirty),实现加锁对性能的影响。
通过引入两个map将读写分离到不同的map，其中read map提供并发读和已存元素原子写，而dirty map则负责读写。 这样read map就可以在不加锁的情况下进行并发读取,当read map中没有读取到值时,再加锁进行后续读取,并累加未命中数,当未命中数大于等于dirty map长度,将dirty map上升为read map。从之前的结构体的定义可以发现，虽然引入了两个map，但是底层数据存储的是指针，指向的是同一份值。

开始时sync.Map写入数据
```
X=1
Y=2
Z=3
```
dirty map主要接受写请求，read map没有数据，此时read map与dirty map数据如下图。
![](https://raw.githubusercontent.com/developer-learning/night-reading-go/master/articles/images/sync-map1.png)

读取数据的时候从read map中读取，此时read map并没有数据，miss记录从read map读取失败的次数，当misses>=len(dirty map)时，将dirty map直接升级为read map,这里直接对dirty map进行地址拷贝并且dirty map被清空，misses置为0。此时read map与dirty map数据如下图。
![](https://raw.githubusercontent.com/developer-learning/night-reading-go/master/articles/images/sync-map2.png)

现在有需求对Z元素进行修改Z=4，sync.Map会直接修改read map的元素。

![](https://raw.githubusercontent.com/developer-learning/night-reading-go/master/articles/images/sync-map3.png)

新加元素K=5，新加的元素就需要操作dirty map了，如果misses达到阀值后dirty map直接升级为read map并且dirty map为空map(read的amended==false)，则dirty map需要从read map复制数据。  

![](https://raw.githubusercontent.com/developer-learning/night-reading-go/master/articles/images/sync-map4.png)

升级后的效果如下。

![](https://raw.githubusercontent.com/developer-learning/night-reading-go/master/articles/images/sync-map5.png)

如果需要删除Z，需要分几种情况：  
一种read map存在该元素且read的amended==false：直接将read中的元素置为nil。
![](https://raw.githubusercontent.com/developer-learning/night-reading-go/master/articles/images/sync-map6.png)

另一种为元素刚刚写入dirty map且未升级为read map:直接调用golang内置函数delete删除dirty map的元素；
![](https://raw.githubusercontent.com/developer-learning/night-reading-go/master/articles/images/sync-map7.png)

还有一种是read map和dirty map同时存在该元素：将read map中的元素置为nil，因为read map和dirty map 使用的均为元素地址，所以均被置为nil。
![](https://raw.githubusercontent.com/developer-learning/night-reading-go/master/articles/images/sync-map8.png)


### 优化点
1. 空间换时间。通过冗余的两个数据结构(read、dirty),实现加锁对性能的影响。
2. 使用只读数据(read)，避免读写冲突。
3. 动态调整，miss次数多了之后，将dirty数据提升为read。
4. double-checking（双重检测）。
5. 延迟删除。 删除一个键值只是打标记，只有在提升dirty的时候才清理删除的数据。
6. 优先从read读取、更新、删除，因为对read的读取不需要锁。



### 方法源码分析

#### Load
Load返回存储在映射中的键值，如果没有值，则返回nil。ok结果指示是否在映射中找到值。  
```go
func (m *Map) Load(key interface{}) (value interface{}, ok bool) {
	// 第一次检测元素是否存在
	read, _ := m.read.Load().(readOnly)
	e, ok := read.m[key]
	if !ok && read.amended {
		// 为dirty map 加锁
		m.mu.Lock()
		// 第二次检测元素是否存在，主要防止在加锁的过程中,dirty map转换成read map,从而导致读取不到数据
		read, _ = m.read.Load().(readOnly)
		e, ok = read.m[key]
		if !ok && read.amended {
			// 从dirty map中获取是为了应对read map中不存在的新元素
			e, ok = m.dirty[key]
			// 不论元素是否存在，均需要记录miss数，以便dirty map升级为read map
			m.missLocked()
		}
		// 解锁
		m.mu.Unlock()
	}
	// 元素不存在直接返回
	if !ok {
		return nil, false
	}
	return e.load()
}
```
dirty map升级为read map
```go
func (m *Map) missLocked() {
	// misses自增1
	m.misses++
	// 判断dirty map是否可以升级为read map
	if m.misses < len(m.dirty) {
		return
	}
	// dirty map升级为read map
	m.read.Store(readOnly{m: m.dirty})
	// dirty map 清空
	m.dirty = nil
	// misses重置为0
	m.misses = 0
}
```
元素取值
```go
func (e *entry) load() (value interface{}, ok bool) {
	p := atomic.LoadPointer(&e.p)
	// 元素不存在或者被删除，则直接返回
	if p == nil || p == expunged {
		return nil, false
	}
	return *(*interface{})(p), true
}
```
read map主要用于读取，每次Load都先从read读取，当read中不存在且amended为true，就从dirty读取数据  。无论dirty map中是否存在该元素，都会执行missLocked函数，该函数将misses+1，当`m.misses < len(m.dirty)`时，便会将dirty复制到read，此时再将dirty置为nil,misses=0。 

#### storage
设置Key=>Value。
```go  
func (m *Map) Store(key, value interface{}) {
	// 如果read存在这个键，并且这个entry没有被标记删除，尝试直接写入,写入成功，则结束
	// 第一次检测
	read, _ := m.read.Load().(readOnly)
	if e, ok := read.m[key]; ok && e.tryStore(&value) {
		return
	}
	// dirty map锁
	m.mu.Lock()
	// 第二次检测
	read, _ = m.read.Load().(readOnly)
	if e, ok := read.m[key]; ok {
		// unexpungelocc确保元素没有被标记为删除
		// 判断元素被标识为删除
		if e.unexpungeLocked() {
			// 这个元素之前被删除了，这意味着有一个非nil的dirty，这个元素不在里面.
			m.dirty[key] = e
		}
		// 更新read map 元素值
		e.storeLocked(&value)
	} else if e, ok := m.dirty[key]; ok {
		// 此时read map没有该元素，但是dirty map有该元素，并需修改dirty map元素值为最新值
		e.storeLocked(&value)
	} else {
		// read.amended==false,说明dirty map为空，需要将read map 复制一份到dirty map
		if !read.amended {
			m.dirtyLocked()
			// 设置read.amended==true，说明dirty map有数据
			m.read.Store(readOnly{m: read.m, amended: true})
		}
		// 设置元素进入dirty map，此时dirty map拥有read map和最新设置的元素
		m.dirty[key] = newEntry(value)
	}
	// 解锁，有人认为锁的范围有点大，假设read map数据很大，那么执行m.dirtyLocked()会耗费花时间较多，完全可以在操作dirty map时才加锁，这样的想法是不对的，因为m.dirtyLocked()中有写入操作
	m.mu.Unlock()
}
```
尝试存储元素。
```go  
func (e *entry) tryStore(i *interface{}) bool {
	// 获取对应Key的元素，判断是否标识为删除
	p := atomic.LoadPointer(&e.p)
	if p == expunged {
		return false
	}
	for {
		// cas尝试写入新元素值
		if atomic.CompareAndSwapPointer(&e.p, p, unsafe.Pointer(i)) {
			return true
		}
		// 判断是否标识为删除
		p = atomic.LoadPointer(&e.p)
		if p == expunged {
			return false
		}
	}
}
```
unexpungelocc确保元素没有被标记为删除。如果这个元素之前被删除了，它必须在未解锁前被添加到dirty map上。
```go  
func (e *entry) unexpungeLocked() (wasExpunged bool) {
	return atomic.CompareAndSwapPointer(&e.p, expunged, nil)
}
```
从read map复制到dirty map。
```go  
func (m *Map) dirtyLocked() {
	if m.dirty != nil {
		return
	}

	read, _ := m.read.Load().(readOnly)
	m.dirty = make(map[interface{}]*entry, len(read.m))
	for k, e := range read.m {
		// 如果标记为nil或者expunged，则不复制到dirty map
		if !e.tryExpungeLocked() {
			m.dirty[k] = e
		}
	}
}
```

#### LoadOrStore
如果对应的元素存在，则返回该元素的值，如果不存在，则将元素写入到sync.Map。如果已加载值，则加载结果为true;如果已存储，则为false。
```go  
func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	// 不加锁的情况下读取read map
	// 第一次检测
	read, _ := m.read.Load().(readOnly)
	if e, ok := read.m[key]; ok {
		// 如果元素存在（是否标识为删除由tryLoadOrStore执行处理），尝试获取该元素已存在的值或者将元素写入
		actual, loaded, ok := e.tryLoadOrStore(value)
		if ok {
			return actual, loaded
		}
	}

	m.mu.Lock()
	// 第二次检测
	// 以下逻辑参看Store
	read, _ = m.read.Load().(readOnly)
	if e, ok := read.m[key]; ok {
		if e.unexpungeLocked() {
			m.dirty[key] = e
		}
		actual, loaded, _ = e.tryLoadOrStore(value)
	} else if e, ok := m.dirty[key]; ok {
		actual, loaded, _ = e.tryLoadOrStore(value)
		m.missLocked()
	} else {
		if !read.amended {
			m.dirtyLocked()
			m.read.Store(readOnly{m: read.m, amended: true})
		}
		m.dirty[key] = newEntry(value)
		actual, loaded = value, false
	}
	m.mu.Unlock()

	return actual, loaded
}
```
如果没有删除元素，tryLoadOrStore将自动加载或存储一个值。如果删除元素，tryLoadOrStore保持条目不变并返回ok= false。
```go  
func (e *entry) tryLoadOrStore(i interface{}) (actual interface{}, loaded, ok bool) {
	p := atomic.LoadPointer(&e.p)
	// 元素标识删除，直接返回
	if p == expunged {
		return nil, false, false
	}
	// 存在该元素真实值，则直接返回原来的元素值
	if p != nil {
		return *(*interface{})(p), true, true
	}

	// 如果p为nil(此处的nil，并是不是指元素的值为nil，而是atomic.LoadPointer(&e.p)为nil，元素的nil在unsafe.Pointer是有值的)，则更新该元素值
	ic := i
	for {
		if atomic.CompareAndSwapPointer(&e.p, nil, unsafe.Pointer(&ic)) {
			return i, false, true
		}
		p = atomic.LoadPointer(&e.p)
		if p == expunged {
			return nil, false, false
		}
		if p != nil {
			return *(*interface{})(p), true, true
		}
	}
}
```

#### Delete
删除元素,采用延迟删除，当read map存在元素时，将元素置为nil，只有在提升dirty的时候才清理删除的数,延迟删除可以避免后续获取删除的元素时候需要加锁。当read map不存在元素时，直接删除dirty map中的元素
```go  
func (m *Map) Delete(key interface{}) {
	// 第一次检测
	read, _ := m.read.Load().(readOnly)
	e, ok := read.m[key]
	if !ok && read.amended {
		m.mu.Lock()
		// 第二次检测
		read, _ = m.read.Load().(readOnly)
		e, ok = read.m[key]
		if !ok && read.amended {
			// 不论dirty map是否存在该元素，都会执行删除
			delete(m.dirty, key)
		}
		m.mu.Unlock()
	}
	if ok {
		// 如果在read中，则将其标记为删除（nil）
		e.delete()
	}
}
```
元素值置为nil
```go  
func (e *entry) delete() (hadValue bool) {
	for {
		p := atomic.LoadPointer(&e.p)
		if p == nil || p == expunged {
			return false
		}
		if atomic.CompareAndSwapPointer(&e.p, p, nil) {
			return true
		}
	}
}
```

#### Range
遍历获取sync.Map中所有的元素，使用的为快照方式，所以不一定是准确的。
```go  
func (m *Map) Range(f func(key, value interface{}) bool) {
	// 第一检测
	read, _ := m.read.Load().(readOnly)
	// read.amended=true,说明dirty map包含所有有效的元素（含新加，不含被删除的），使用dirty map
	if read.amended {
		// 第二检测
		m.mu.Lock()
		read, _ = m.read.Load().(readOnly)
		if read.amended {
			// 使用dirty map并且升级为read map
			read = readOnly{m: m.dirty}
			m.read.Store(read)
			m.dirty = nil
			m.misses = 0
		}
		m.mu.Unlock()
	}
	// 一贯原则，使用read map作为读
	for k, e := range read.m {
		v, ok := e.load()
		// 被删除的不计入
		if !ok {
			continue
		}
		// 函数返回false，终止
		if !f(k, v) {
			break
		}
	}
}
```

### 总结
经过了上面的分析可以得到,sync.Map并不适合同时存在大量读写的场景,大量的写会导致read map读取不到数据从而加锁进行进一步读取,同时dirty map不断升级为read map。 从而导致整体性能较低,特别是针对cache场景.针对append-only以及大量读,少量写场景使用sync.Map则相对比较合适。

sync.Map没有提供获取元素个数的Len()方法，不过可以通过Range()实现。
```go
func Len(sm sync.Map) int {
	lengh := 0
	f := func(key, value interface{}) bool {
		lengh++
		return true
	}
	one:=lengh
	lengh=0
	sm.Range(f)
	if one != lengh {
	    one = lengh
		lengh=0
		sm.Range(f)
		if one <lengh {
			return lengh
		}
		
	}
	return one
}
```

---
参考
* [Go sync.Map](https://github.com/golang/go/blob/master/src/sync/map.go)
* [Go 1.9 sync.Map揭秘](http://colobu.com/2017/07/11/dive-into-sync-Map/)