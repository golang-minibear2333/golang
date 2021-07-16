# 4.2 channel

## **...本节正在编写，未完待续，催更请留言，我会收到邮件**

> 本节源码位置 https://github.com/golang-minibear2333/golang/blob/master/4.concurrent/channel.go

我们平时肯定没少接触过队列，在

channel 是 `goroutine` 之间互相通讯的东西。类似我们 `Unix` 上的管道（可以在进程间传递消息），用来 `goroutine` 之间发消息和接收消息。其实，就是在做 `goroutine` 之间的内存共享。

`channel`是类型相关的，也就是说一个 `channel` 只能传递一种类型的值，这个类型需要在 `channel` 声明时指定。

## 4.2.1 声明与初始化

channel 的一般声明形式：

```go
var chanName chan 类型
```

与普通变量的声明不同的是在类型前面加了 `channel` 关键字，`类型` 则指定了这个 `channel` 所能传递的元素类型。示例：

```go
var a chan int //声明一个传递元素类型为int的channel
var b chan float64
var c chan string
```

初始化一个 `channel` 也非常简单，直接使用 `Go` 语言内置的 `make()` 函数，示例：

```go
a := make(chan int) //初始化一个int型的名为a的channel
b := make(chan float64)
c := make(chan string)
```

channel 最频繁的操作就是写入和读取，这两个操作也非常简单，示例：

```go
a := make(chan int)
a <- 1  //将数据写入channel
z := <-a  //从channel中读取数据
fmt.Println(z)
```

正常来说，一个发送，一个接收，协作合理，完美无缺。但是，事实令人咋舌。输出死锁：

```go
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        .../golang-minibear2333/golang/4.concurrent/channel.go:7 +0x59
```

来看下完整的代码

```go
func main() {
    a := make(chan int)
    a <- 1   //将数据写入channel
    z := <-a //从channel中读取数据
    fmt.Println(z)
}
```

* 观察上面三行代码，第 2 行往 channel 内写入了数据，第 3 行从 channel 中读取了数据
* 但是这是在一个方法中，并且没有使用 Go 关键字，说明他们在同一个协程
* 

如果程序运行正常当然不会出什么问题，可如果第二行数据写入失败，或者 channel 中没有数据，那么第 3 行代码会因为永远无法从 a
中读取到数据而一直处于阻塞状态。

## 4.4.1 超时机制



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

## 小结

这一节简单介绍了go语言中的channel（信道），go语言主张不要通过共享内存来通信，而应通过通信来共享内存，通过`channel`的方式可以完成不同`goroutine`之间的通信。