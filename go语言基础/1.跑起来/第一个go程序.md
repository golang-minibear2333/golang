### 第一个 go 程序

带着目标学东西往往是最有成效的，**为什么学**以及**环境安装**可以参考前面的文章。
> go语言github项目：https://github.com/golang-minibear2333/golang


let's go go go !

```go
package main

import "fmt"

func main() {
   fmt.Println("Hello, World!")
}
```

这是一个最简单的 go 程序，由三个元素构成

**元素一、包** ：`package`是包，每个文件夹是一个包，默认包名就是文件夹名，文件夹下所有的`.go`文件全部都是同一个包。

`main`包整个项目的入口，你可以在指定任意一个文件夹当作程序的入口，然后把里面所有文件第一行改成`package main`，也可以把项目根目录当作`main`包。

PS: 每个项目默认有且只有一个`main`包

**元素二、import**：这个关键字代表引入其他地方的包，可以是当前项目的，也可以是别人写的。
这里`import "fmt"`引入的就是 go 原生的`fmt`包，专门用来输出文本的。

**元素三、语句**

```go
func main() {
   fmt.Println("Hello, World!")
}
```

main 入口函数，每个项目只有一个`main`函数

**执行以上代码输出**

```go
$ go run hello.go
Hello, World!
```

得，今天就是这么短，如果你觉得这点知识太特么少了，请告诉我。


