goroutine 是 Go 语言并行设计的核心。goroutine 是一种比线程更轻量的实现，十几个 goroutine 可能在底层就是几个线程。 不同的是，Golang 在 runtime、系统调用等多方面对 goroutine
调度进行了封装和处理，当遇到长时间执行或者进行系统调用时，会主动把当前 goroutine 的 CPU (P) 转让出去，让其他 goroutine 能被调度并执行，也就是 Golang 从语言层面支持了协程。要使用 goroutine
只需要简单的在需要执行的函数前添加 go 关键字即可。当执行 goroutine 时候，Go 语言立即返回，接着执行剩余的代码，goroutine 不阻塞主线程。下面我们通过一小段代码来讲解 go 的使用：

```go
//首先我们先实现一个 Add()函数
func Add(a, b int) {
c := a + b
fmt.Println(c)
}

go Add(1, 2) //使用go关键字让函数并发执行
```

Go 的并发执行就是这么简单，当在一个函数前加上 go 关键字，该函数就会在一个新的 goroutine 中并发执行，当该函数执行完毕时，这个新的 goroutine
也就结束了。不过需要注意的是，如果该函数具有返回值，那么返回值会被丢弃。所以什么时候用 go 还需要酌情考虑。

接着我们通过一个案例来体验一下 Go 的并发到底是怎么样的。新建源文件 [goroutine.go](goroutine.go)，输入以下代码：

```go
package main

import "fmt"

func Add(a, b int) {
	c := a + b
	fmt.Println(c)
}

func main() {
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}
}
```

执行 goroutine.go 文件会发现屏幕上什么都没有，但程序并不会报错，这是什么原因呢？原来当主程序执行到 for 循环时启动了 10 个 goroutine，然后主程序就退出了，而启动的 10 个 goroutine 还没来得及执行
Add() 函数，所以程序不会有任何输出。也就是说主 goroutine 并不会等待其他 goroutine 执行结束。那么如何解决这个问题呢？Go 语言提供的信道（channel）就是专门解决并发通信问题的，下一节我们将详细介绍。