---
title: "timeout"
date: 2021-06-26T02:39:47+08:00
draft: true
---

# timeout

## 超时机制

通过前面的内容我们了解到，channel 的读写操作非常简单，只需要通过 <- 操作符即可实现，但是 channel 的使用不当却会带来大麻烦。我们先来看之前的一段代码：

```go
a := make(chan int)
a <- 1
z := <-a
```

观察上面三行代码，第 2 行往 channel 内写入了数据，第 3 行从 channel 中读取了数据，如果程序运行正常当然不会出什么问题，可如果第二行数据写入失败，或者 channel 中没有数据，那么第 3 行代码会因为永远无法从 a
中读取到数据而一直处于阻塞状态。

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

## channel 的关闭

channel 的关闭非常简单，使用 Go 语言内置的 close() 函数即可关闭 channel，示例：

```go
ch := make(chan int)
close(ch)
```

关闭了 channel 后如何查看 channel 是否关闭成功了呢？很简单，我们可以在读取 channel 时采用多重返回值的方式，示例：

```go
x, ok := <-ch
```

通过查看第二个返回值的 bool 值即可判断 channel 是否关闭，若为 false 则表示 channel 被关闭，反之则没有关闭。

建议使用 defer 如下

```go
t := make(chan bool)
ch := make(chan int)
defer func() {
    close(ch)
    close(t)
}()
```

[完整代码](timeout.go)

```go
package main

import "time"

func main() {
	t := make(chan bool)
	ch := make(chan int)
	defer func() {
		close(ch)
		close(t)
	}()
	go func() {
		time.Sleep(1e9) //等待1秒
		t <- true
	}()
	go func() {
		time.Sleep(time.Second * 2)
		ch <- 123
	}()
	select {
	case <-ch: //从ch中读取数据

	case <-t: //如果1秒后没有从ch中读取到数据，那么从t中读取，并进行下一步操作
	}
}
```