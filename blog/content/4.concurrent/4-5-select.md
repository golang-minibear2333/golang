# 4.4 select

## **...本节正在编写，未完待续，催更请留言，我会收到邮件**

> 本节源码位置 https://github.com/golang-minibear2333/golang/blob/master/4.concurrent/4.5-select

## 4.3.1 select与switch

让我们来复习一下`switch`语句，在`switch`语句中，会逐个匹配`case`语句(可以是值也可以是表达式)，一个一个的判断过去，直到有符合的语句存在，执行匹配的语句内容后跳出`switch`。

```go
func demo(number int){
    switch{
        case number >= 90:
        fmt.Println("优秀")
        default:
        fmt.Println("太搓了")
    }
}
```

而 `select` 用于处理通道，它的语法与 `switch` 非常类似。

```golang
func main() {
	chanInt1, chanInt2 := make(chan int), make(chan int)
	go func() {
		defer close(chanInt1)
		defer close(chanInt2)
		chanInt1 <- 1
		chanInt2 <- 2
	}()
	time.Sleep(time.Millisecond)
	select {
	case data := <-chanInt1:
		fmt.Println(data)
	case data := <-chanInt2:
		fmt.Println(data)
	default:
		fmt.Println("全部阻塞")
	}
}
```

比如这样接收数值



## 4.3.4 超时机制

相反的，如果 channel 中的数据一直没有被读取，那么写入操作也会一直处于阻塞状态。如果不正确处理这个情况，很可能会导致整个 goroutine 锁死，这就是超时问题。Go
语言没有针对超时提供专门的处理机制，但是我们却可以利用 select 来巧妙地实现超时处理机制，下面看一个示例：

```go
t := make(chan bool)
go func() {
    time.Sleep(1e9) //等待1秒
    t <- true
}()

select {
    case <-ch:  //从ch中读取数据

    case <-t:  //如果1秒后没有从ch中读取到数据，那么从t中读取，并进行下一步操作
}
```

这样的方法就可以让程序在等待 1 秒后继续执行，而不会因为 ch 读取等待而导致程序停滞，从而巧妙地实现了超时处理机制，这种方法不仅简单，在实际项目开发中也是非常实用的。

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

## 泄露防止

及时通知select

## 小结

