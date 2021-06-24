> go语言github项目：https://github.com/golang-minibear2333/golang

[TOC]

### 映射关系容器 map

`Go`语言提供的映射关系容器为 `map` ， `map` 使用散列表`hash`实现。查找复杂度为O(1)，和数组一样，最坏的情况下为O(n),n为元素总数。

这就是`Go`中`map`的定义格式。
```go
map[keyType] valueType
```
注意了，map 是一种引用类型，初值是`nil`,定义时必须用`make`来创建，否则会报错 
```
panic: assignment to entry in nil map
```
必须要申请空间，所有的引用类型都要这么做
```go
var m map[string]string
m = make(map[string]string) 
```
当然，也可以这么写
```go
m := make(map[string]string) 
```

###  使用

赋值
```go
	m["name"] = "coding3min"
	m["sex"] = "man"
```

循环遍历
```go
for key := range m {
    // 原来不用Printf也可以完成拼接输出啊！
		fmt.Println("key:", key, ",value:", m[key]) 
}
```

删除集合元素
```go
	delete(m, "name")
```

PS: 在取值的时候`m[key]`，假如`key`不存在，不会报错，会返回`value`类型的默认值，比如`int`类型默认值为`0`

当然了，如果你想明确的知道元素是否存在，如下：

```go
	if value, ok := m[key]; ok {
		fmt.Println(key, "存在，值为：", value)
	} else {
		fmt.Println(key, " 不存在")
	}
```

map容器就到这里了。

### 能够在并发环境中使用的`map`

`Go`中的`map`在并发读的时候没问题，但是并发写就不行了（线程不安全），会发生竞态问题。

所以有一个叫`sync.Map`的封装数据结构供大家使用，简单用法如下：
定义和存储
```go
	var scene sync.Map
	scene.Store("name", "coding3min")
	scene.Store("age", 11)
```
取值
```go
v, ok := scene.Load("name")
	if ok {
		fmt.Println(v)
	}
	v, ok = scene.Load("age")
	if ok {
		fmt.Println(v)
	}
```

输出

```
coding3min
11
```

删除和遍历,这里遍历就用到了[函数当作参数传递](https://mp.weixin.qq.com/s/HsaEjO9TgUcfrBhaMS0C5A)和[匿名函数](https://mp.weixin.qq.com/s/YRD2-4oO9ENHD3ADYlvsCg)的知识。

```go
	scene.Delete("age")

	scene.Range(func(key, value interface{}) bool {
		fmt.Println("key:",key,",value:",value)
		return true
	})
```

the end，今天的更新还没完，次条还有切片知识补充哦

