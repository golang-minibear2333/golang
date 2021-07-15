# 4.3 select

## **...本节正在编写，未完待续，催更请留言，我会收到邮件**

> 本节源码位置 https://github.com/golang-minibear2333/golang/blob/master/4.concurrent/select.go

## 4.3.1 select与switch

让我们来复习一下`switch`语句，在`switch`语句中，会逐个匹配`case`语句，一个一个的判断过去，直到有符合的语句存在，执行匹配的语句内容后跳出`switch`。

内部可以是值，也可以是表达式，如果`switch`后未接参数，就必须是已有变量的表达式。

```go
switch number{
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

而 `select` 用于处理异步 `IO` 问题，它的语法与 `switch` 非常类似。

## 4.3.2 select

由 select 开始一个新的选择块，每个选择条件由 case 语句来描述，并且每个 case 语句里必须是一个 channel 操作。它既可以用于 channel
的数据接收，也可以用于 channel 的数据发送。如果 select 的多个分支都满足条件，则会随机的选取其中一个满足条件的分支。

新建源文件 [channel.go](channel.go)，输入以下代码：

```go
func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()
	start := time.Now() // 获取当前时间

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	elapsed := time.Since(start)
	// 这里没有用到3秒，为什么？
	fmt.Println("该函数执行完成耗时：", elapsed)
}
```

以上代码先初始化两个 channel c1 和 c2，然后开启两个 goroutine 分别往 c1 和 c2 写入数据，再通过 select 监听两个 channel，从中读取数据并输出。

运行结果如下：
```shell
$ go run channel.go 
received one
received two
该函数执行完成耗时： 2.004695535s
```

