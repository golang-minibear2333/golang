# 1.4 声明【变量】的各种方式

> 本节源码位置 https://github.com/golang-minibear2333/golang/blob/master/1.base/1.4-variables

讲变量就要先知道 go 语言有哪些数据类型。

## 1.4.1 数据类型

数据类型的出现是为了把数据分成所需内存大小不同的数据。

- 布尔型(`bool`): 值只可以是常量 `true` 或者 `false`。
- 数字类型: 整型 `int` 和浮点型 `float`，支持复数（业务代码用不到），其中位的运算采用补码。
- 字符串类型（`string`）: 使用`UTF-8`编码标识`Unicode`文本。
- 其他：指针、数组、结构体(`struct`)、联合体 (`union`不常用)、函数、切片、接口（`interface`）、`Map` 、 `Channel`

大多数类型都是接触过的，比如`c++`的结构体，比如`python`的切片，`java`的接口，别看类型那么多以后写多了自然就会用了。

go 语言声明变量的方式非常简单

## 1.4.2 第一种方式、var

```go
var name string
```

结构为`var`+`变量名`+`类型`

```go
name = "s"
```

像这样赋值

```go
//根据赋值自动判断类型
var p = name
```

因为`name`是字符串类型，所以`p`也是同类型

```go
//多变量声明,int类型不赋值自动赋值为0，比如d e f
var a, b, c = 1, 2, 3
var d, e, f int
```

一次声明多个类型不同的变量

```go
//类型不同的多个变量，难看的要死
	var (
		k int
		l string
	)

//这样好看
var m, n, o = "a", 1, true
```

## 1.4.3 方式二、:=

```go
//直接声明并赋值（必须是初次声明才有冒号）
p2 := "as"
// 多个变量一次性声明并赋值
h, i, j := 1, 2, 3
```

## 1.4.4 常量

常量就是**不可变的变量**，定义方式

```go
const identifier [type] = value
```

约定常量全大写表示

```go
const A int = 1
const B = 1
const C, D, E = 1, 1, 1
```

一般常量被用于枚举

```go
	const (
		Success = 0
		UnKonw  = 1
		Error   = 2
	)
```

不过要枚举还是用 `go` 自带的特殊常量好一点，这种特殊被认为是可以被编译器修改的常量

```go
	//const 出现时被重置为0，每出现一次自动加1
	const (
		F = iota
		G = iota
		H = iota
	)
```

F、G、H 值为0，1，2

当然可以简写成这样，效果是一样的。

```go
	const (
		I = iota
		J
		K
	)
```

## 1.4.5 类型转换

没有什么好说的，和其他语言相似，类型转换都是类型+变量的形式，如下。

```go
 var aInt int = 17

	// 一般用这种方式强制转
	fmt.Printf("转float64 %f  \n", float64(aInt))
	fmt.Printf("转string %v  \n", strconv.Itoa(aInt))
	fmt.Printf("转float64 %f  \n", float64(aInt))[]
```

输出

```
转float64 17.000000  
转string 17  
转float64 17.000000  
```

各种类型转字符串

```go
	resString := fmt.Sprintf("%d %v %v", 1, "coding3min", true)
	fmt.Println(resString)
```

输出

```go
1 coding3min true
```

`string`  和 `bytes` 的互相转换

```go
// string  to bytes
	resBytes := []byte("asdfasdf")
	// bytes to string
	resString = string(resBytes)
	fmt.Println(resString)
```

输出

```
asdfasdf
```

## 1.4.6 小结

本节介绍了常量和变量，以及变量之间简单类型的转换，这里语言的基础，需要熟练掌握，特别是在做算法的时候更是高频用到。