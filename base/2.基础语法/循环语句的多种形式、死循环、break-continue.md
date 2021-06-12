
今天 go 语言的内容是循环。

> go语言github项目：https://github.com/golang-minibear2333/golang

[TOC]

由于在不少实际问题中有许多具有**规律性的重复操作**，因此在程序中就需要重复执行某些语句。

`go` 语言的循环和其他的没什么不同，只是语法上略微有些差别。

### for 循环方式 1 和 c++、java 相似

```go
nums := []int{1, 2, 3, 4, 5, 6}

for i := 0; i < len(nums); i++ {
		fmt.Println(i)
	}
```

### for 循环方式 2 省略赋值和++

```go
a, b := 1, 5
for a < b {
		fmt.Println(a)
		a++
	}
```

### for 循环方式 3 迭代

- 优点：不用引入无意义的变量
- 缺点：不是直接索引，如果数据量极大会有性能损耗

```go
for index, value := range nums {
		fmt.Printf("key: %v , value: %v
    \n", index, value)
	}
```

当然，你可以把方式 3 中 `index` 去掉,用`_`忽略掉`key`

```go
	for _, v := range nums {
		fmt.Printf("value: %v \n", v)
	}
```

如果你想忽略掉 `value`，直接用 `key`也是可以的，这样就消除了迭代方式的缺点！

```go
	for i := range nums {
		fmt.Printf("value: %v \n", nums[i])
	}
```

### 死循环
这样就是一个最简单的死循环,循环条件永远为`true`也是死循环
```go
for {
}
```
### break、continue

i := 0
```go
	for {
		fmt.Printf("死循环测试 %v \n", i)
		i++
		if i > 5 {
			fmt.Println("满足终止条件，退出")
			break //直接跳出循环
		}
		if i == 3 {
			continue //会直接跳过执行后面的语句
		}
		fmt.Printf("死循环测试,第%v次跑到循环结尾\n", i)
	}
```
输出
```go
死循环测试 0 
死循环测试,第1次跑到循环结尾
死循环测试 1 
死循环测试,第2次跑到循环结尾
死循环测试 2 
死循环测试 3 
死循环测试,第4次跑到循环结尾
死循环测试 4 
死循环测试,第5次跑到循环结尾
死循环测试 5 
满足终止条件，退出
```


