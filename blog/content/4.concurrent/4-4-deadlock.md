# 4.4 deadlock

> 本节源码位置 https://github.com/golang-minibear2333/golang/blob/master/4.concurrent/4.4-deadlock/

## 4.4.1 什么时候会导致死锁

在计算机组成原理里说过 死锁有三个必要条件他们分别是 **循环等待、资源共享、非抢占式**，在并发中出现通道死锁只有两种情况：

- 数据要发送，但是没有人接收
- 数据要接收，但是没有人发送

## 4.4.2 发送单个值时的死锁

牢记这两点问题就很清晰了，复习下之前的例子，会死锁

```go
a := make(chan int)
a <- 1   //将数据写入channel
z := <-a //从channel中读取数据
```

- 有且只有一个协程时，无缓冲的通道
- 无论是先发送会阻塞在发送，先接收会阻塞在接收处。
- 因为发送操作在接收者准备好之前是阻塞的，接收操作在发送之前是阻塞的，
- 解决办法就是改为缓冲通道，或者使用协程配对

解决方法一,协程配对，先发送还是先接收无所谓只要配对就好

```go
chanInt := make(chan int)
go func() {
    chanInt <- 1
}()

res := <-chanInt
```

解决方法二，缓冲通道

```go
chanInt := make(chan int,1)
chanInt <- 2
res := <-chanInt
```

- 缓冲通道内部的消息数量用`len()`函数可以测试出来
- 缓冲通道的容量可以用`cap()`测试出来
- 在满足`cap>len`时候，因为没有满，发送不会阻塞
- 在`len>0`时，因为不为空，所以接收不会阻塞

使用缓冲通道可以让生产者和消费者减少阻塞的可能性，对异步操作更友好，不用等待对方准备，但是容量不应设置过大，不然会占用较多内存。

## 4.4.3 多个值发送的死锁

配对可以让死锁消失，但发送多个值的时候又无法配对了，又会死锁

```go
func multipleDeathLock() {
	chanInt := make(chan int)
	defer close(chanInt)
    go func() {
		res := <-chanInt
		fmt.Println(res)
	}()
	chanInt <- 1
	chanInt <- 1
}
```

不出所料死锁了

```go
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.multipleDeathLock()
```

在工作中通知信号是一对一的情况，通知一次以后就不再使用了，其他这种要求多次读写配对的情况根本不会存在。

## 4.4.4 解决多值发送死锁

更常见的是用循环来不断接收值，接受一个处理一个，如下：

```go
func multipleLoop() {
	chanInt := make(chan int)
	defer close(chanInt)
	go func() {
		for {
			//不使用ok会goroutine泄漏
			//res := <-chanInt
			res,ok := <-chanInt
			if !ok {
                 break
            }
			fmt.Println(res)
		}
	}()
	chanInt <- 1
	chanInt <- 1
}
```

输出：

```go
1
1
```

- 给通道的接收加上二值，`ok` 代表通道是否正常，如果是关闭则为`false`值
- 可以删掉那段逻辑试试，会输出`1 2 0 0 0`这样的数列，因为关闭是需要时间的，而循环接收关闭的通道拿到的是`0`
- 关于`goroutine`泄漏稍后会讲到

## 4.4.5 应该先发送还是先接收

假如我们调换一下位置，把接收放外面，写入放里面会发生什么

```go
func multipleDeathLock2() {
	chanInt := make(chan int)
	defer close(chanInt)
	go func() {
		chanInt <- 1
		chanInt <- 2
	}()
	for {
		res, ok := <-chanInt
		if !ok {
			break
		}
		fmt.Println(res)
	}
}
```

输出死锁

```go
1
2
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.multipleDeathLock2()
```

- 出现上面的结果是因为`for`循环一直在获取通道中的值，但是在读取完`1 2`后，通道中没有新的值传入，这样接收者就阻塞了。
- 为什么先接收再发送可以，因为发送提前结束后会触发函数的`defer`自动关闭通道
- 所以我们应该总是先接收后发送，并由发送端来关闭

## 4.4.6 goroutine 泄漏

`goroutine` 终止的场景有三个：

- 当一个 `goroutine` 完成了它的工作
- 由于发生了没有处理的错误
- 有其他的协程告诉它终止

当三个条件都没有满足，`goroutine` 就会一直运行下去

```go
func goroutineLeak() {
	chanInt := make(chan int)
	defer close(chanInt)
	go func() {
		for {
            res := <-chanInt
			//res,ok := <-chanInt
			//if !ok {
            //     break
            //}
			fmt.Println(res)
		}
	}()
	chanInt <- 1
	chanInt <- 1
}
```

- 上面的`goroutineLeak()`函数结束后触发`defer close(chanInt)`关闭了通道
- 但是匿名函数中`goroutine`并没有关闭，而是一直在循环取值，并且取到是的关闭后的通道值（这里是`int`的默认值 0）
- `goroutine`会永远运行下去，如果以后再次使用又会出现新的泄漏！导致内存、`cpu`占用越来越多

输出，如果程序不停止就会一直输出`0`

```go
1
1
0
0
0
...
```

假如不关闭且外部没有写入值，那接收处就会永远阻塞在那里，连输出都不会有

```go
func goroutineLeakNoClosed() {
	chanInt := make(chan int)
	go func() {
		for {
            res := <-chanInt
			fmt.Println(res)
		}
	}()
}
```

- 无任何输出的阻塞
- 换成写入也是一样的
- 如果是有缓冲的通道，换成已满的通道写没有读；或者换成向空的通道读没有写也是同样的情况
- 除了阻塞，`goroutine`进入死循环也是泄露的原因

## 4.4.7 如何发现泄露

使用 golang 自带的`pprof`监控工具，可以发现内存上涨情况，这个后续会讲

还可以监控进程的内存使用情况，比如`prometheus`提供的`process-exporter`

如果你有内存泄露/goroutine 泄露代码扫描的工具，欢迎留言，感恩！

## 4.4.8 小结

今天我们学习了一些细节，但是相当重要的知识点，也是未来面试高频问题哦！

- 如果是信号通知，应该保证一一对应，不然会死锁
- 除了信号通知外，通常我们使用循环处理通道，在工作中不断的处理数据
- 应该总是先接收后发送，并由发送端来关闭，不然容易死锁或者泄露
- 在接收处，应该对通道是否关闭做好判断，已关闭应该退出接收，不然会泄露
- 小心 goroutine 泄漏，应该在通道关闭的时候及时检查通道并退出
- 除了阻塞，`goroutine`进入死循环也是泄露的原因
