# 4.5 select

> 本节源码位置 https://github.com/golang-minibear2333/golang/blob/master/4.concurrent/4.5-select

## 4.5.1 select与switch

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

而 `select` 用于处理通道，它的语法与 `switch` 非常类似。每个 `case` 语句里必须是一个 `channel` 操作。它既可以用于 `channel` 的数据接收，也可以用于 `channel` 的数据发送。

```golang
func foo() {
	chanInt := make(chan int)
	defer close(chanInt)
	go func() {
		select {
		case data, ok := <-chanInt:
			if ok {
				fmt.Println(data)
			}
		default:
			fmt.Println("全部阻塞")
		}
	}()
    time.Sleep(time.Second)
    chanInt <- 1
}
```

输出`1`

* 这是一个简单的接收发送模型。
* 如果 select 的多个分支都满足条件，则会随机的选取其中一个满足条件的分支。
* 第6行加上ok是因为上一节讲过，如果不加会导致通道关闭时收到零值。
* 回忆之前的知识，接收和发送应该在不同的`goroutine`里。
* 其次`select default`子协程，在`case`都处于阻塞状态时，会直接执行`default`的内容。导致子协程提前退出，主协程中的写入操作会一直阻塞(等待接收者，接收者已经退出了) 触发死锁
* 倒数第二行加了`sleep` 1秒，是因为让`select`语句提前结束的问题暴露出来。

```go
全部阻塞
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.bar()
```

`select` 执行完了，退出了`goroutine`，而发送才刚刚执行到，没有与其匹配的接收，故死锁。

正确的做法是把接收套在循环里面。

```go
func bar() {
	chanInt := make(chan int)
	defer close(chanInt)
	go func() {
		for {
			select {
			    ...
			}
		}
	}()
	chanInt <- 1
}
```

* 不再死锁了
* 假如程序不停止，会出现一个泄露的`goroutine`，永远的在`for`循环中无法跳出，此时引入下一节的内容

## 4.5.2 通知机制

`Go` 语言总是简单和灵活的，虽然没有针对提供专门的机制来处理退出，但我们可以自己组合

```go
func main() {
	chanInt, done := make(chan int), make(chan struct{})
	defer close(chanInt)
	defer close(done)
	go func() {
		for {
			select {
			case <-chanInt:
			case <-done:
				return
			}
		}
	}()
	done <- struct{}{}
}
```

* 没有给`chanInt`发送任何东西，按理说会阻塞，导致`goroutine`泄露
* 但可以使用额外的通道完成协程的退出控制
* 这种方式还可以做到周期性处理任务，下一节我们再详细讲解

## 4.5.3 case执行原理

假如`case`后左边和右边跟了函数，会执行函数，我们来探索一下。

定义`A`、`B`函数，作用相同

```go
func A() int {
	fmt.Println("start A")
	time.Sleep(1 * time.Second)
	fmt.Println("end A")
	return 1
}
```

定义函数`lee`，请问该函数执行完成耗时多少呢？

```go
func lee() {
	ch, done := make(chan int), make(chan struct{})
	defer close(ch)
	go func() {
		select {
		case ch <- A():
		case ch <- B():
		case <-done:
		}
	}()
	done <- struct{}{}
}
```

答案是2秒

```go
start A
end A
start B
end B
main.leespend time: 2.003504395s
```

* select扫描是从左到右从上到下的，按这个顺序先求值，如果是函数会先执行函数。
* 然后立马判断是否可以立即执行（这里是指case是否会因为执行而阻塞）。
* 所以两个函数都会进入，而且是先进入A再进入B，两个函数都会执行完，所以等待时间会累计。
* 所以不应该在case判断中放函数。

如果都不会阻塞，此时就会使用一个伪随机的算法，去选中一个case，只要选中了其他就被放弃了。

## 4.5.4 超时控制

我们来模拟一个更真实点的例子，让程序一段时间超时退出。

定义一个结构体

```go
type Worker struct {
	stream  <-chan int //处理
	timeout time.Duration //超时
	done    chan struct{} //结束信号
}
```

定义初始化函数

```go
func NewWorker(stream <-chan int, timeout int) *Worker {
	return &Worker{
		stream:  stream,
		timeout: time.Duration(timeout) * time.Second,
		done:    make(chan struct{}),
	}
}
```

定义超时处理函数

```go
func (w *Worker) afterTimeStop() {
	go func() {
		time.Sleep(w.timeout)
		w.done <- struct{}{}
	}()
}
```

* 超过时间发送结束信号

接收数据并处理函数

```go
func (w *Worker) Start() {
	w.afterTimeStop()
	for {
		select {
		case data, ok := <-w.stream:
			if !ok {
				return
			}
			fmt.Println(data)
		case <-w.done:
			close(w.done)
			return
		}
	}
}
```

* 收到结束信号关闭函数
* 这样的方法就可以让程序在等待 1 秒后继续执行，而不会因为 ch 读取等待而导致程序停滞。
  
```go
func main() {
	stream := make(chan int)
	defer close(stream)

	w := NewWorker(stream, 3)
	w.Start()
}
```

实际3秒到程序运行结束。好在官方已经考虑到这一点，为我们提供了现成的方案。

## 4.5.5 官方超时方案

```go
go func() {
    t := time.NewTicker(timeout)
    defer t.Stop()
    for {
        select {
        case data := <-chanInt:
            t.Reset(timeout)
        case <-t.C:
        case <-done:
            return
        }
    }
}()
```

* `time.NewTicker`创建了一个定时器，参数为时间间隔，并返回一个结构体`t`。
* `t.C` 是一个仅可接收的`channel`，会根据时间间隔定时执行任务，也可以作为超时任务使用。
* `t.Reset(timeout)` 重置时间，因为`select`进入一个`case`，后续的执行会有耗时，所以要重置时间保证时间的精准。

这种方式巧妙地实现了超时处理机制，这种方法不仅简单，在实际项目开发中也是非常实用的。

在生产中，常常把`buf`积累到一定数量然后`flush`出去，假如数据产生速度太慢，就要靠定时器定时消费，看下面完整的例子。

```go
func main() {
	chanInt, done := make(chan int), make(chan struct{})
	defer close(chanInt)
	defer close(done)
	go func() {
		...
	}()
	for i := 0; i < 100; i++ {
		if i%10 == 0 {
			time.Sleep(time.Second)
		}
		chanInt <- 1
	}
	done <- struct{}{}
}
```

产生100个数，每10个数暂停1秒，用来模拟数据产生速度慢，`go func()` 内容如下：

```go
go func() {
    timeout := time.Second
    t := time.NewTicker(timeout)
    defer t.Stop()
    buf := make([]int, 0, 5)
    for {
        select {
        case data := <-chanInt:
            t.Reset(timeout)
            if len(buf) < cap(buf) {
                buf = append(buf, data)
            } else {
                go send(buf)
                buf = make([]int, 0, cap(buf))
            }
        case <-t.C:
            if len(buf) > 0 {
                go send(buf)
                buf = make([]int, 0, cap(buf))
            }
        case <-done:
            return
        }
    }
}()
```

* 接收到数据时，如果`buf`满了就进行上报，如果`buf`没满就追加数据。
* 假如超时，就直接发送`buf`防止数据太少一直不发送的情况。
* 需要在其他case里，`Reset`超时时间，以校准定时器。

## 4.5.6 小结

本节介绍了`select`的用法以及包含的陷阱，我们学会了：

* `case`只针对通道传输阻塞做特殊处理，如果有计算将会先进行计算，所以不应该在case判断中放函数。
* 扫描是从左到右从上到下的，按这个顺序先求值，如果是函数会先执行函数。如果函数运行时间长，时间会累计。
* 在`case`全部阻塞时，会执行`default`中的内容。
* 可使用结束信号，让`select`退出。
* 延时发送结束信号可以实现超时自动退出的功能。
* 官方的`time`包，提供了定时器，可作定时任务，也可作超时控制。

我还写了可热更新的定时器，有兴趣了解的可以看看本节的源码哦。