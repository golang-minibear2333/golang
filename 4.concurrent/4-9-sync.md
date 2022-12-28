# 4.9 sync包

> 本节源码位置 https://github.com/golang-minibear2333/golang/tree/master/4.concurrent/4.9-sync/

## 4.9.1 sync.Map 并发安全的Map

反例如下，两个`Goroutine`分别读写。

```go
func unsafeMap(){
	var wg sync.WaitGroup
	m := make(map[int]int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			m[i] = i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			fmt.Println(m[i])
		}
	}()
	wg.Wait()
}
```

执行报错：

```go
0
fatal error: concurrent map read and map write

goroutine 7 [running]:
runtime.throw({0x10a76fa, 0x0})
......
```

使用并发安全的`Map`

```go
func safeMap() {
	var wg sync.WaitGroup
	var m sync.Map
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			m.Store(i, i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			fmt.Println(m.Load(i))
		}
	}()
	wg.Wait()
}
```

* 不需要`make`就能使用
* 还内置了`Store`、`Load`、`LoadOrStore`、`Delete`、`Range`等操作方法，自行体验。


## 4.9.2 sync.Once 只执行一次

很多场景下我们需要确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等。

`init` 函数是当所在的 `package` 首次被加载时执行，若迟迟未被使用，则既浪费了内存，又延长了程序加载时间。

`sync.Once` 可以在代码的任意位置初始化和调用，因此可以延迟到使用时再执行，并发场景下是线程安全的。

在多数情况下，`sync.Once` 被用于控制变量的初始化，这个变量的读写满足如下三个条件：

* 当且仅当第一次访问某个变量时，进行初始化（写）；
* 变量初始化过程中，所有读都被阻塞，直到初始化完成；
* 变量仅初始化一次，初始化完成后驻留在内存里。

```go
var loadOnce sync.Once
var x int
for i:=0;i<10;i++{
    loadOnce.Do(func() {
        x++
    })
}
fmt.Println(x)
```

输出1

## 4.9.3 sync.Cond 条件变量控制

`sync.Cond` 基于互斥锁/读写锁，它和互斥锁的区别是什么呢？

互斥锁 `sync.Mutex` 通常用来保护临界区和共享资源，条件变量 `sync.Cond` 用来协调想要访问共享资源的 `goroutine`。

也就是在存在共享变量时，可以直接使用`sync.Cond`来协调共享变量，比如最常见的共享队列，多消费多生产的模式。

我一开始也很疑惑为什么不使用`channel`和`select`的模式来做生产者消费者模型（实际上也可以），这一节不是重点就不展开讨论了。

创建实例

```go
func NewCond(l Locker) *Cond
```

`NewCond` 创建 `Cond` 实例时，需要关联一个锁。

广播唤醒所有

```go
func (c *Cond) Broadcast()
```

`Broadcast` 唤醒所有等待条件变量 `c` 的 `goroutine`，无需锁保护。

唤醒一个协程

```go
func (c *Cond) Signal()
```

`Signal` 只唤醒任意 1 个等待条件变量 `c` 的 `goroutine`，无需锁保护。

等待

```go
func (c *Cond) Wait()
```

每个 Cond 实例都会关联一个锁 L（互斥锁 *Mutex，或读写锁 *RWMutex），当修改条件或者调用 Wait 方法时，必须加锁。

举个不恰当的例子，实现一个经典的生产者和消费者模式，但有先决条件：

* 边生产边消费，可以多生产多消费。
* 生产后通知消费。
* 队列为空时，暂停等待。
* 支持关闭，关闭后等待消费结束。
* 关闭后依然可以生产，但无法消费了。

```go
var (
	cnt          int
	shuttingDown = false
	cond         = sync.NewCond(&sync.Mutex{})
)
```

* `cnt` 为队列，这里直接用变量代替了，变量就是队列长度。
* `shuttingDown` 消费关闭状态。
* `cond` 现成的队列控制。

生产者

```go
func Add(entry int) {
	cond.L.Lock()
	defer cond.L.Unlock()
	cnt += entry
	fmt.Println("生产咯，来消费吧")
	cond.Signal()
}
```

消费者

```go
func Get() (int, bool) {
	cond.L.Lock()
	defer cond.L.Unlock()
	for cnt == 0 && !shuttingDown {
		fmt.Println("未关闭但空了，等待生产")
		cond.Wait()
	}
	if cnt == 0 {
		fmt.Println("关闭咯，也消费完咯")
		return 0, true
	}
	cnt--
	return 1, false
}
```

关闭程序

```go
func Shutdown() {
	cond.L.Lock()
	defer cond.L.Unlock()
	shuttingDown = true
	fmt.Println("要关闭咯，大家快消费")
	cond.Broadcast()
}
```

主程序

```go
var wg sync.WaitGroup
	wg.Add(2)
	time.Sleep(time.Second)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			go Add(1)
			if i%5 == 0 {
				time.Sleep(time.Second)
			}
		}
	}()
	go func() {
		defer wg.Done()
		shuttingDown := false
		for !shuttingDown {
			var cur int
			cur, shuttingDown = Get()
			fmt.Printf("当前消费 %d, 队列剩余 %d \n", cur, cnt)
		}
	}()
	time.Sleep(time.Second * 5)
	Shutdown()
	wg.Wait()
```

* 分别创建生产者与消费者。
* 生产10个，每5个休息1秒。
* 持续消费。
* 主程序关闭队列。

输出

```go
生产咯，来消费吧
当前消费 1, 队列剩余 0 
未关闭但空了，等待生产
生产咯，来消费吧
生产咯，来消费吧
当前消费 1, 队列剩余 1 
当前消费 1, 队列剩余 0 
未关闭但空了，等待生产
生产咯，来消费吧
生产咯，来消费吧
生产咯，来消费吧
当前消费 1, 队列剩余 2 
当前消费 1, 队列剩余 1 
当前消费 1, 队列剩余 0 
未关闭但空了，等待生产
生产咯，来消费吧
生产咯，来消费吧
生产咯，来消费吧
生产咯，来消费吧
当前消费 1, 队列剩余 1 
当前消费 1, 队列剩余 2 
当前消费 1, 队列剩余 1 
当前消费 1, 队列剩余 0 
未关闭但空了，等待生产
要关闭咯，大家快消费
关闭咯，也消费完咯
当前消费 0, 队列剩余 0
```

## 4.9.4 小结

* sync.Map 并发安全的Map。
* sync.Once 只执行一次，适用于配置读取、通道关闭。
* sync.Cond 控制协调共享资源。


## 引用

* [Go sync.Once](https://geektutu.com/post/hpg-sync-once.html)
* [k8s workerqueue源码](https://sourcegraph.com/github.com/kubernetes/client-go@b69cda3e4e09e1b956d42709fbadf12e6c0f24b7/-/blob/util/workqueue/queue_test.go?L269:5-269:22)