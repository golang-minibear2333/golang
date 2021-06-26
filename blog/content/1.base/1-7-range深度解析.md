# 1.7 range深度解析

> 本节源码位置 https://github.com/golang-minibear2333/golang/blob/master/1.base/1.7-range-deep

## 1.7.1 range（范围）

`range` 关键字在 `go` 语言中是相当常用好用的语法糖，可以用在 `for` 循环中迭代 `array`、`slice`、`map`、`channel`、`字符串`所有涉及到遍历输出的东西。

## 1.7.2 怎么用？

我们在前一节 [循环](https://mp.weixin.qq.com/s/hxeysXVCPKR7Wlql9D7YlA) 中初次触及到了 `range`，也知道他可以省略`key`，或者省略`value`来循环遍历的特性，但是这种特性要结合实际情况考量该用哪一个。

**切片迭代**

```go
	nums := []int{1, 2, 3}
	for k, v := range nums {
		fmt.Printf("key: %v , value: %v  \n", k, v)
	}
```

这和迭代方式非常适合用`for-range`语句，如果减少赋值，直接索引`num[key]`可以减少损耗。如下

```go
for k, _:= range nums {
```

**`map`迭代**
注意，从 `Go1`开始，遍历的起始节点就是随机了。

```go
	kvs := map[string]string{
		"a":"a",
		"b":"b",
	}
	for k, v := range kvs {
		fmt.Printf("key: %v , value: %v  \n", k, v)
	}
```

函数中`for-range`语句中只获取 `key` 值，然后跟据 `key` 值获取 `value` 值，虽然看似减少了一次赋值，但通过 `key` 值查找 `value` 值的性能消耗可能高于赋值消耗。

所以能否优化取决于 map 所存储数据结构特征、结合实际情况进行。

**字符串迭代**(一个一个的输出字符)

```go
for k,v := range "hello"{
  //注意这里单个字符输出的是ASCII码，
  //用 %c 代表输出字符
		fmt.Printf("key: %v , value: %c \n", k, v)
	}
```

**channel** （如果不会可以先 mark 下，详细参考后续：go 的并发特性章节）

```go
ch := make(chan int, 10)
	ch <- 11
	ch <- 12

	close(ch) // 不用的时候记得关掉,不关掉又没有另一个goroutine存在会死锁哦，可以注释掉这一句体验死锁

	for x := range ch {
		fmt.Println(x)
	}
```
**结构体**
```go
tmp := []struct{
		int
		string
	}{
		{1, "a"},
		{2, "b"},
	}

	for k,v := range tmp{
		fmt.Printf("k:%v, v:%v  \n",k,v)
	}
```

**注意**：由于循环开始前循环次数就已经确定了，所以循环过程中新添加的元素是没办法遍历到的。

## 1.7.3 有可能会遇到的坑！

由于`range`遍历时`value`是值的拷贝，如果这个时候遍历上一节声明的结构体时，修改`value`，原结构体不会发生任何变化！
```go
for _,v := range tmp{
		v.a = 2
	}
```
两次输出一致
```go
k:0, v:{1 a}  
k:1, v:{2 b}  
k:0, v:{1 a}  
k:1, v:{2 b}  
```

## 1.7.4 编程 Tips

- 遍历过程中可以适情况放弃接收 `index` 或 `value`，可以一定程度上提升性能
- 遍历 `channel` 时，如果 `channel` 中没有数据，可能会阻塞
- 尽量避免遍历过程中修改原数据

## 1.7.5 小结

* `range`可以用于`for` 循环中迭代 `array`、`slice`、`map`、`channel`、`字符串`所有涉及到遍历输出的东西。
* `for-range` 的实现实际上是`C`风格的`for`循环
* 使用`index`,`value`接收`range`返回值会发生一次数据拷贝

