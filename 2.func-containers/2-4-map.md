# 2.4 map

> 本节源码位置 https://github.com/golang-minibear2333/golang/blob/master/2.func-containers/2.4-map

## 2.4.1 映射关系容器 map

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

## 2.4.2 使用

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

## 2.4.3 map 内部元素的修改

map 可以拷贝吗？

`map` 其实是不能拷贝的，如果想要拷贝一个 `map` ，只有一种办法就是循环赋值，就像这样

```go
originalMap := make(map[string]int)
originalMap["one"] = 1
originalMap["two"] = 2

// Create the target map
targetMap := make(map[string]int)

// Copy from the original map to the target map
for key, value := range originalMap {
    targetMap[key] = value
}
```

如果 `map` 中有指针，还要考虑深拷贝的过程

```go
originalMap := make(map[string]*int)
var num int = 1
originalMap["one"] = &num

// Create the target map
targetMap := make(map[string]*int)

// Copy from the original map to the target map
for key, value := range originalMap {
var tmpNum int = *value
    targetMap[key] = &tmpNum
}
```

如果想要更新 `map` 中的`value`，可以通过赋值来进行操作

```go
map["one"] = 1
```

但如果 `value` 是一个结构体，可以直接替换结构体，但无法更新结构体内部的值

```go
originalMap := make(map[string]Person)
originalMap["minibear2333"] = Person{age: 26}
originalMap["minibear2333"].age = 5
```

你可以 [试下源码函数 updateMapValue](https://github.com/golang-minibear2333/golang/blob/master/2.func-containers/2.4-map/map1.go#L89) ，会报这个错误

> Cannot assign to originalMap["minibear2333"].age

问题链接 [issue-3117](https://github.com/golang/go/issues/3117) , 其中 [ianlancetaylor](https://github.com/golang/go/issues/3117#issuecomment-430632750) 的回答很好的解释了这一点

简单来说就是map不是一个并发安全的结构，所以，并不能修改他在结构体中的值。

这如果目前的形式不能修改的话，就面临两种选择，

* 1.修改原来的设计; 
*  2.想办法让map中的成员变量可以修改，

因为懒得该这个结构体，就选择了方法2

要么创建个临时变量，做拷贝，像这样

```go
tmp := m["foo"]
tmp.x = 4
m["foo"] = tmp
```

要么直接用指针，比较方便

```go
originalPointMap := make(map[string]*Person)
originalPointMap["minibear2333"] = &Person{age: 26}
originalPointMap["minibear2333"].age = 5
```

## 2.4.4 能够在并发环境中使用的`map`

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

删除和遍历,这里遍历就用到了 [函数当作参数传递](https://mp.weixin.qq.com/s/HsaEjO9TgUcfrBhaMS0C5A) 和 [匿名函数](https://mp.weixin.qq.com/s/YRD2-4oO9ENHD3ADYlvsCg) 的知识。

```go
	scene.Delete("age")

	scene.Range(func(key, value interface{}) bool {
		fmt.Println("key:",key,",value:",value)
		return true
	})
```

## 2.4.5 小结

本节介绍了字典`map`类型，这种类型在很多语言中都有，并且学习了它的增加删除元素的方法，以及更新value要注意的点。

还介绍了并发环境下使用的线程安全的 `sync.Map`。

