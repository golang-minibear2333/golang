# 4.6

很多时候需要周期性的执行某些操作，就需要用到定时器。定时器有三种思路。

## 4.6.1 Sleep

使用休眠，让当前`Goroutine`休眠一定的时间来实现定时的效果，缺点是程序执行速度不均匀，导致定时周期不均匀。

```go
for{
fmt.Println(time.Now())
time.Sleep(time.Second*1)
}
```

## 4.6.2 Timer

`Go` 语言的内置包，指定一个时间开始计时，时间到之后会向外发送通知，发送通知的方式就是使用`<-chan Time` 返回内容。

第一种方式，直接在需要等待处使用，效果和`Sleep`一样，一使用就卡在那了内部就是使用了`Timer`。

```go
    fmt.Println(time.Now())
<-time.After(1*time.Second)
fmt.Println(time.Now())
```

也可以把他拆分开，在任意地方进行等待

```go
    timer := time.NewTimer(1 * time.Second)
<-timer.C
fmt.Println(time.Now())
```

但是以上只是做到延迟一次性执行，我们来改造一下，把他变成定时器。

```go
    done := make(chan struct{})
timer := time.NewTimer(1 * time.Second)
go func () {
for {
select {
case <-timer.C:
fmt.Println(time.Now())
timer.Reset(1 * time.Second)
case <-done:
return
}
}
}()
<-time.After(5*time.Second + time.Millisecond*100)
done <- struct{}{}
```

* 定义子`Goroutine`的目的是为了防止形成死锁，让定时器最终能退出，在实际项目中可能需要一个永久运行的定时器，一般为了不影响项目主逻辑也会这样定义。如果你的项目就是定时任务，我建议也这么写，这样可以注册很多个定时器互不影响。
* `done`是为了判断执行是否结束，防止主`Goroutine`提前退出。
* 这个示例只有两个`case`，实战中如果有加其他`case`需要给每个`case`内都做一次`Reset`，保证重置定时器。

## 4.6.3 Ticker

相比上述使用延迟执行功能实现的定时器，`Ticker` 本身就是一个定时器(内部封装了`Timer`)，我们使用起来就非常简单。

```go
ticker := time.NewTicker(1 * time.Second)
go func () {
for {
<-ticker.C
fmt.Println(time.Now())
}
}()
<-time.After(5 * time.Second + time.Millisecond*100)
ticker.Stop()
```

在[select](https://golang.coding3min.com/4.concurrent/4-5-select/) 一节中讲述的官方超时控制方案非常的实用，也是使用的此函数。还使用到`timer.Stop`
和`timer.Reset`这两个内置函数这里就不展开讲解了，建议进行复习。

## 4.6.4 小结

定时器一般用来周期性执行任务，比如定时同步数据、计算报表、发送通知。

* `time.Sleep` 使用休眠，让当前`goroutine`休眠一定的时间来实现定时的效果，缺点是内部逻辑执行的速度会影响到定时器的时间差，无法做到精确间隔。
* `Timer` 类似于`Sleep`的延迟处理，通过`channel`来获得通知，也可以改造成定时器。因为是延迟处理，所以要记得重置时间来实现定时执行的效果。
* `Ticker` 现成的定时器，内部也是封装了 `Timer`。