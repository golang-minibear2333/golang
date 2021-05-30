### 简介

`Goroutine` 是 `Golang` 中非常有用的功能，有时候 `goroutine` 没执行完函数就返回了，如果希望等待当前的 `goroutine` 执行完成再接着往下执行，该怎么办？

``` go
package main

import (
    "time"
    "fmt"
)

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

### 唯一好方案

Golang 官方在 sync 包中提供了 WaitGroup 类型来解决这个问题。其文档描述如下：

> A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.

大意为： `WaitGroup` 用来等待单个或多个 `goroutines` 执行结束。在主逻辑中使用 `WaitGroup` 的 `Add` 方法设置需要等待的 `goroutines` 的数量。在每个 goroutine 执行的函数中，需要调用 `WaitGroup` 的 `Done` 方法。最后在主逻辑中调用 `WaitGroup` 的 `Wait` 方法进行阻塞等待，直到所有 `goroutine` 执行完成。
使用方法可以总结为下面几点：

* 创建一个 `WaitGroup` 实例，比如名称为：wg
* 调用 `wg.Add(n)` ，其中 n 是等待的 `goroutine` 的数量
* 在每个 `goroutine` 运行的函数中执行 `defer wg.Done()`
* 调用 `wg.Wait()` 阻塞主逻辑

``` Go
package main

import (
    "time"
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    say2("hello", &wg)
    say2("world", &wg)
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

简短的例子，注意循环传入的变量用中间变量替代，防止闭包 `bug`

``` Go
var wg sync.WaitGroup
	wg.Add(len(sList))
	for _, d := range sList {
		tmpD := d
		go func(waitGroup *sync.WaitGroup) {
            defer waitGroup.Done()
            // to do something
            // use tmpD 
            }
		}(&wg)
	}
	wg.Wait()
```

### 引用 

[Golang 入门 : 等待 goroutine 完成任务](https://www.cnblogs.com/sparkdev/p/10917536.html)
