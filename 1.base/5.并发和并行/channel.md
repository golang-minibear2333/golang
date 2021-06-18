channel 是goroutine 之间互相通讯的东西。类似我们 Unix 上的管道（可以在进程间传递消息），用来 goroutine 之间发消息和接收消息。其实，就是在做 goroutine 之间的内存共享。channel
是类型相关的，也就是说一个 channel 只能传递一种类型的值，这个类型需要在 channel 声明时指定。

### 声明与初始化

channel 的一般声明形式：var chanName chan ElementType。

与普通变量的声明不同的是在类型前面加了 channel 关键字，ElementType 则指定了这个 channel 所能传递的元素类型。示例：

```go
var a chan int //声明一个传递元素类型为int的channel
var b chan float64
var c chan string
```

初始化一个 channel 也非常简单，直接使用 Go 语言内置的 make() 函数，示例：

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
```