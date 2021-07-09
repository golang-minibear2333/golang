# 2.1 函数简单使用和基本知识解析

> 本节源码位置 https://github.com/golang-minibear2333/golang/blob/master/2.func-containers/2.1-func

这里的源码有多个，本节相关的有

```shell
# 函数的简单使用
main.go
# 函数当作变量使用，当做 参数传递
function_value.go
# 值传递和引用传递
more.go
# 递归函数
recursive.go
```

拓展代码有

```shell
# 函数当作变量使用，当做 参数传递的分页实践
function_value_good_demo.go
# 函数方法(go中定义一个类)
go_class.go
```

## 2.1.1 基本原理

函数，几乎是每种编程语言的必备语法，通过函数把一系列的动作汇总起来，在不同的地方重复使用。

我们在数学中曾经就使用过函数，他的形式类似于`y=f（x）`，这就是一个完整的调用过程，`y`就是函数计算后得到的值，`x`就是传入的变量。

## 2.1.2 怎么用？

相信在看这个教程的人肯定已经接触过其他的编程语言，我就不多废话了，就是干。

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-04-25-152325.jpg)

go语言中最基本的函数是这样的，以`func`为关键字标记函数

```go
func functionParam(num int) {
}
```

当然了，可以有多个形参，类型相同时可以省略，如下

```go
//多个参数的函数
func functionParams(a, b int, c string) {
}
```

上面说过的函数都没有返回值，一般的函数都有返回值，没有返回值的函数要么是**引用传递**，可以直接改变参数内容，要么就是用于**单元测试**或者打印输出等。

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-04-25-151012.jpg)

没有返回值的函数就像一个不完整的男人，只能接受不能输出，来看看这个男人有一个输出的情况。

返回值标记在函数第一个括号后面，由于`go`语言是强类型语言，但又和`python`不同，要写出返回值类型。

```go
//一个返回值
func funcReturnOne() int {
	return 1
}
```

如果说是有多个返回值，要用打括号括起来。

```go
//多个返回值
func funReturnMany() (int, int) {
	return 1, 2
}
```

上面的返回值全部都是匿名的，可以赐他一个名字，函数中不用定义返回值，可以省略几行代码。

```go
//返回值有名称
func funReturnName() (res int) {
  //var res  int  省掉了
	res = 1 + 1
	return
}
```

用返回就有接收，函数外部用这种方式接收

```go
	//接收多个返回值
	a, b := funReturnMany()
```

## 2.1.3 值传递，引用传递

刚刚有说到函数没有返回值的时候，要么是只需要打印结果，要么是只做单元测试，除了这两种情况，没有返回值的函数就是做了很多事情的你**没有和老板汇报**一样，没有任何意义！

引用传递和`c++`类似，先举个值传递的例子。

```go
//值传递
func noChange(a, b int) {
	tmp := a
	a = b
	b = tmp
}
```

调用打印结果看看 

```go
  a, b := 1, 2
	fmt.Printf("原值 a:%v,b:%v \n", a, b)
	noChange(a, b)
	//值传递，并没有修改原值
	fmt.Printf("值传递后 a:%v,b:%v \n", a, b)
```

看！像不像任劳任怨的你，忙活半天被老板喜欢的小张抢了功劳。

```
原值 a:1,b:2 
值传递后 a:1,b:2 
```

下面来看看引用传递的例子。在类型前加一个星号代表该参数是一个指针

```go
// 引用传递，参数加*号代表指针
func change(a,b *int){
	tmp := *a
	*a = *b
	*b = tmp
}
```

学过`c++`再来学`go`简直是如虎添翼，`c++`中有一个指针的概念`go`语言里也有。

```go
	//引用传递，&就是c中的取地址
	change(&a,&b)
	fmt.Printf("引用传递后 a:%v,b:%v \n", a, b)
```

输出结果，可以看到值被调换了。引用传递需要加`&`符号，术语叫`取地址`。函数里的对他做的任何操作都会改变原来的变量内容。 

```go
引用传递后 a:2,b:1 
```

上面的例子传入的是指针，还有一种叫引用类型，和指针的区别是不需要星号和`&`，对他的修改会直接改动到原有变量的值。

ps:go语言中只有三种引用类型，slice(切片)、map(字典)、channel(管道)

## 2.1.4 函数进阶

上面说的东西都很简单了，基本学过任何一门语言的人都能瞬间看懂，和`python`、`c++`、`javascript`一样，`go`中也有把函数当作参数传递的语法。

像这样，`functionValue`函数的形参里有一个名为`do`的函数，需要提前指定`do`函数有什么参数和返回值。

```go
func functionValue(a, b int, do func(int, int) int) {
	fmt.Println(do(a, b))
}
```

然后`do(a,b)`是在`functionValue`内部调用的。这种特性有什么用呢？定义两个参数为`int`，返回为`int`的函数。

```go
func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
```

因为规则符合`do`函数的规则，两个都可以传递过去，看！这就不用修改函数内部而出现了两种效果。

```go
	functionValue(1, 1, add)
	functionValue(1, 1, sub)
```

在设计模式里，这种方式叫**装饰器模式（Decorator Pattern）**:允许向一个现有的对象添加新的功能，同时又不改变其结构。

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-04-25-155351.jpg)

当然，你也不必每次传递函数的时候都憨厚老实的定义一个新函数，因为有时候你定义的函数就只会在这里用到，只不过是把实现放在调用外部，而不修改原函数代码罢了。

```go
	//匿名函数
	functionValue(1, 1, func(i1 int, i2 int) int {
		return i1 * i2
	})
```

上面这个例子多看几遍啊！！

## 2.1.5 实际的使用

你可以参考函数测速例子

定义一个测速函数。

```go
func speedTime(handler func() (string), funcName string) {
	t := time.Now()
	handler()
	elapsed := time.Since(t)
	fmt.Println(funcName+"spend time:", elapsed)
}
```

传入不同的函数都可以测速度。

```go
speedTime(appendStr, "appendStr")
speedTime(appendStrQuick, "appendStrQuick")
```

小`Tips`：
- 还有你可以传filter函数做过滤，mapping做映射等实际的用法
- 有时候也可以作为排序递增，递减的依据

## 2.1.6 小结

本节讲述了Go中函数的基本语法，包括定义、多值返回，函数的值传递和引用传递，还可以当变量来用，可以把函数当参数来传递
