# 4.3 channel

到这里你正在接触最核心和重要的知识！认真学习的你很棒！

> 本节源码位置 https://github.com/golang-minibear2333/golang/blob/master/4.concurrent/4.3-channel/

## 4.3.1 什么是 channel

Go 是一门从语言级别就支持并发的编程语言， 它有一个设计哲学很特别 **不要通过共享内存来通信，而应通过通信来共享内存** ，听起来是有一点绕。

在传统语言中并发使用全局变量来进行不同线程之间的数据共享，这种方式就是使用共享内存的方式进行通信。而 Go 会在协程和协程之间打一个隧道，通过这个隧道来传输数据（发送和接收）。

![4-3-channel](https://coding3min.oss-accelerate.aliyuncs.com/2021/07/19/4-3-channel.png)

打个比方，我们平时肯定没少接触过队列，队列的特点是先进先出，多方生产插入，多方消费接收。这个队列/隧道就是`channel`。

`channel` 是 `goroutine` 之间互相通讯的东西，`goroutine` 之间用来发消息和接收消息。其实，就是在做 `goroutine` 之间的内存共享。

我们来看看具体是什么使用的。

## 4.3.2 声明与初始化

`channel`是类型相关的，也就是说一个 `channel` 只能传递一种类型的值，这个类型需要在 `channel` 声明时指定。

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

通道是一个引用类型，初始值为`nil`，对于值为`nil`的通道，不论具体是什么类型，它们所属的接收和发送操作都会永久处于阻塞状态。

所以必须手动`make`初始化，示例：

```go
a := make(chan int) //初始化一个int型的名为a的channel
b := make(chan float64)
c := make(chan string)
```

既然是队列，那就有大小，上面没声明具体的大小，被认为是*无缓冲*的（注意大小是 0，不是 1）也就是说必须有其他`goroutine`接收，不然就会阻塞在那。声明有缓冲的，指定大小就可以了。

```go
a := make(chan int,100)
```

## 4.3.3 如何使用

我们进一步体验一下无缓冲 channel 会发生什么问题，同时熟悉下用法，示例：

```go
func pendingForever() {
a := make(chan int)
a <- 1   //将数据写入channel
z := <-a //从channel中读取数据
fmt.Println(z)
}
```

- 观察上面三行代码，第 2 行往 `channel` 内写入了数据，第 3 行从 `channel` 中读取了数据
- 但是这是在同一个方法中，并且没有使用 Go 关键字，说明他们在同一个协程

我们说过 `channel` 是用来给不同 `goroutine` 通信的，所以是不能在同一个协程又发送又接收，这根本就达不到隧道通信的效果。所以上面的代码，会死锁：

```go
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
.../4.concurrent/channel.go:7 +0x59
```

死锁的原因是没有其他协程来接收数据，隧道因为是无缓冲的，所以直接永远的阻塞在发送方。

要解决这个问题也好办。放到不同 `goroutine` 里就可以。

```go
func normal() {
chanInt := make(chan int)
go func() {
chanInt <- 1
}()

res := <-chanInt
fmt.Println(res)
}
```

输出`1`。无缓冲通道在无数据发送时，接收端会阻塞，直到有新数据发送过来为止。

上面的代码，一个发送一个接收，而实际使用中数据往往是连续不断发送的。来看一段代码：

```go
func standard() {
chanInt := make(chan int)
go func() {
defer close(chanInt)
var produceData = []int{1, 2, 3}
for _, v := range produceData {
chanInt <- v
}
}()
for v := range chanInt {
fmt.Println(v)
}
}
```

输出

```go
1
2
3
```

- 循环传递数据，父协程循环接收。
- `range chan` 的方式可以不断的接收数据，直到通道关闭，假如通道不关闭会永远阻塞，无法通过编译，直接报死锁。
- 必须在发送端关闭通道，因为接收端无法预料是否还有数据没有接收完；向已关闭的`channel`发送数据会`panic`。
- 建议使用 `defer` 来关闭通道，防止程序异常时未正常关闭。

至此我们完成了一个简单的生产者消费者模型。

## 4.3.4 channel 的关闭

使用 Go 语言内置的 `close()` 函数即可关闭 `channel`，再强调一次建议使用`defer`关闭，示例：

```go
defer close(ch)
```

关闭了 `channel` 后如何查看 `channel` 是否关闭成功了呢？很简单，我们可以在读取 `channel` 时采用多重返回值的方式，示例：

```go
x, ok := <-ch
```

通过查看第二个返回值的 `bool` 值即可判断 `channel` 是否关闭，若为 `false` 则表示 `channel` 被关闭，反之则没有关闭（使用频率不高，了解即可）

```go
func main() {
var chanInt chan int = make(chan int, 10)
go func() {
defer fmt.Println("chanInt is closed")
defer close(chanInt)
chanInt <- 1
}()
res := <-chanInt
fmt.Println(res)
}
```

输出

```go
chanInt is closed
1
```

- 如上声明了一个有缓冲的通道，在缓冲大小允许的范围内不需要阻塞等待接收
- 发送端发送完毕后主动关闭通道
- 虽然通道已经关闭，接收端依然可以接收，接收完自行结束。

PS1: 同一个通道只能关闭一次，重复关闭会`panic`。

PS2: 如果传入`nil`,如 `close(nil)` 会 `panic`。

## 4.3.5 多发送、多接收与单向通道

我们结合前面知识，来实战练习一下！

功能：实现一个多发送，多接收的例子。

```go
func send(c chan<- int, wg *sync.WaitGroup) {
c <- rand.Int()
wg.Done()
}
```

- 发送端随机生成数字，并声明一个仅发送的单向通道
- 使用`sync.WaitGroup`做等待（忘记的回顾上一节哈！）

```go
func received(c <-chan int, wg *sync.WaitGroup) {
for gotData := range c {
fmt.Println(gotData)
}
wg.Done()
}
```

- 接收端使用`range`来接收数字并打印

```go
func main() {
chanInt := make(chan int, 10)
done := make(chan struct{})
defer close(done)
go func() {
defer close(chanInt)
// 发送
}()
go func() {
...
// 接收
done <- struct{}{}
}()
<-done
}
```

- 使用了两个通道，一个通道`chanInt`进行数据传输，另一个`done`控制完毕时结束主协程
- 发送端负责生产数据，生产完毕后关闭通道
- 接收端负责接收完毕后通知主协程

发送端

```go
go func() {
    var wg sync.WaitGroup
    defer close(chanInt)
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go send(chanInt, &wg)
    }
    wg.Wait()
}()
```

连续启动 5 个协程，使用`wg`做协程等待，发送完毕再结束是为了交给`defer`关闭`chanInt`

接收端

```go
go func() {
    var wg sync.WaitGroup
    for i := 0; i < 8; i++ {
        wg.Add(1)
        go received(chanInt, &wg)
    }
    wg.Wait()
    done <- struct{}{}
}()
```

连续启动多个接收端，通道被关闭时纷纷退出，最后通知`done`

输出 5 个随机数，程序正常关闭。

```go
5577006791947779410
8674665223082153551
4037200794235010051
6129484611666145821
3916589616287113937
```

单向通道限制了函数的使用方式，它可以用在循环比较耗时的场景，处理完一个数据立马发送出来，尽量减少内存的使用。

## 4.3.6 小结

这一节简单介绍了 go 语言中的 channel（信道），go 语言主张不要通过共享内存来通信，而应通过通信来共享内存，通过`channel`的方式可以完成不同`goroutine`之间的通信。

我们学会了：

- `channel` 是引用类型默认值是`nil`，需要手动`make`。
- 通道必须在多个`goroutine`中使用
- 有缓冲与无缓冲通道的特点，什么时候会阻塞。
- 可以用`range`来做循环接收，通道关闭会自动停止。
- 只能且必须在发送端使用`defer`关闭通道。
- 正式使用一般多发送多接收，并使用`done`信号通知的方式进行通知。

在工作中，通道的使用更为复杂，下一节将介绍`select`，敬请期待！
