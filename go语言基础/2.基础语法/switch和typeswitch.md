> go语言github项目：https://github.com/minibear2333/how_to_code
![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-04-14-125225.jpg)

[TOC]

### ifelse
![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-04-14-125215.jpg)

```go
if 20>0{
		fmt.Println("yes")
	}
```
输出
```go
yes
```

ifelse

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-04-14-125252.jpg)

```go
if 20<0{

	}else{
		fmt.Println("no")
	}
```
输出
```go
no
```

### switch 和  type switch

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-04-14-125322.jpg)

`switch` 好理解,是一个替代`if else else else`接口而提出的,如下，`switch` 后跟变量，`case` 后跟常量，只要变量值和常量匹配，就执行该分支下的语句。
```go
switch name {
	case "coding3min":
		fmt.Println("welcome" + name)
	default:
		fmt.Println("403 forbidden:" + name)
		return
	}
```

当然`switch`语句会逐个匹配`case`语句，一个一个的判断过去，直到有符合的语句存在。
```go
switch {
	case number >= 90:
		fmt.Println("优秀")
	case number >= 80:
		fmt.Println("良好")
	case number >= 60:
		fmt.Println("凑合")
	default:
		fmt.Println("太搓了")
	}
```
如果没有一个是匹配的，就执行`default`后的语句。

注意`switch`后可以跟空，如上
```go
switch {
```
这样`case`就必须是表达式。

### switch 的高级玩法？

有一个流传于坊间的神秘玩法，可以用`switch`语句来判断传入变量的类型，然后做一些羞羞的事情。`x`是一个未知类型的变量，`switch t := x.(type)` 用这个方式来赋值，`t`就是有确定类型的变量。

```go
switch t := x.(type) {
	case int:
		return t
	case float64:
		return int(math.Ceil(t))
	}
```

什么叫未知类型？？

这就是 `go` 中有意思的地方了, `interface{}` 类型，是一种神奇的类型，他可以是任何类型的接口，而具体的类型是实现。
```go
var x interface{}
	x = 1
	fmt.Println(x)
```
输出`1`

所以完整的函数是这样的
```go
func typeSwitchDemo(x interface{}) int {
	switch t := x.(type) {
	case int:
		return t
	case float64:
		return int(math.Ceil(t))
	}
	return 0
}
```
这个东西有什么用呢？？有没有想过如果你有一个场景，你在调用第三方的接口，却发现对方的接口发生了微调，原来的`int`类型，被转换成了`string`类型，你必须写出兼容两种方式的代码来解析`json`。

那么这个时候，`type switch` 将会是你的武器。

感兴趣可以 跑到这里看看，我是怎么使用这个武器的。

https://github.com/minibear2333/how_to_code/blob/master/golang/medium/json_interface/fixed_json.go


