# 4.2 并发等待

> 本节源码位置 https://github.com/golang-minibear2333/golang/blob/master/4.concurrent/goroutine-wait/

## 4.2.1 简介

`goroutine` 是 `Golang` 中非常有用的功能，有时候 `goroutine` 没执行完函数就返回了，如果希望等待当前的 `goroutine` 执行完成再接着往下执行，该怎么办？

```go
func say(s string) {
    for i := 0; i < 3; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("hello world")
    fmt.Println("over!")
}
```

输出 `over！` , 主线程没有等待

## 4.2.2 使用 Sleep 等待

```Go
func main() {
    go say("hello world")
    time.Sleep(time.Second*1)
    fmt.Println("over!")
}
```

运行修改后的程序，结果如下：

```go
hello world
hello world
hello world
over!
```

结果符合预期，但是太 low 了，我们不知道实际执行中应该等待多长时间，所以不能接受这个方案！

## 4.2.3 发送信号

```go
func main() {
    done := make(chan bool)
    go func() {
        for i := 0; i < 3; i++ {
            time.Sleep(100 * time.Millisecond)
            fmt.Println("hello world")
        }
        done <- true
    }()

    <-done
    fmt.Println("over!")
}
```

输出的结果和上面相同，也符合预期

这种方式不能处理多个协程，所以也不是优雅的解决方式。

## 4.2.4 WaitGroup

Golang 官方在 sync 包中提供了 WaitGroup 类型可以解决这个问题。其文档描述如下：

使用方法可以总结为下面几点：

- 在父协程中创建一个 `WaitGroup` 实例，比如名称为：wg
- 调用 `wg.Add(n)` ，其中 n 是等待的 `goroutine` 的数量
- 在每个 `goroutine` 运行的函数中执行 `defer wg.Done()`
- 调用 `wg.Wait()` 阻塞主逻辑
- 直到所有 `goroutine` 执行完成。

```Go
func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    go say2("hello", &wg)
    go say2("world", &wg)
    fmt.Println("over!")
    wg.Wait()
}

func say2(s string, waitGroup *sync.WaitGroup) {
    defer waitGroup.Done()

    for i := 0; i < 3; i++ {
        fmt.Println(s)
    }
}
```

输出，注意顺序混乱是因为并发执行

```go
hello
hello
hello
over!
world
world
world
```

## 4.2.5 小心缺陷

简短的例子，注意循环传入的变量用中间变量替代，防止闭包 `bug`

```Go
func errFunc() {
	var wg sync.WaitGroup
	sList := []string{"a", "b"}
	wg.Add(len(sList))
	for _, d := range sList {
		go func() {
			defer wg.Done()
			fmt.Println(d)
		}()
	}
	wg.Wait()
}
```

输出，可以发现全部变成了最后一个

```go
b
b
```

父协程与子协程是并发的。父协程上的`for`循环瞬间执行完了，内部的协程使用的是`d`最后的值，这就是闭包问题。

解决方法当作参数传入

```go
func correctFunc() {
	var wg sync.WaitGroup
	sList := []string{"a", "b"}
	wg.Add(len(sList))
	for _, d := range sList {
		go func(str string) {
			defer wg.Done()
			fmt.Println(str)
		}(d)
	}
	wg.Wait()
}
```

输出

```go
b
a
```

要留意 `range` 中的`value`有可能出现 **1.7.3 有可能会遇到的坑！**

## 引用

[Golang 入门 : 等待 goroutine 完成任务](https://www.cnblogs.com/sparkdev/p/10917536.html)
