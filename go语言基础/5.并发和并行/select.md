select 用于处理异步 IO 问题，它的语法与 switch 非常类似。由 select 开始一个新的选择块，每个选择条件由 case 语句来描述，并且每个 case 语句里必须是一个 channel 操作。它既可以用于 channel
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
$ go run channel.go received one received two
```
