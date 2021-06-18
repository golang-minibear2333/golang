
### 前言

> go语言github项目：https://github.com/golang-minibear2333/golang

[TOC]

每次新建项目，不熟悉go的项目结构，一般跑都跑不起来，每次都要重新搞一遍，好几回跑项目都会报类似`File is  invalid`的错误

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-03-08-145230.jpg)

有时候报其他奇怪的错误，今天就下决心整理一下，理一理概念 GOROOT、GOPATH、src、 pkg、bin，希望以后不要再出现这样的问题了，同时给看到文章的你一些帮助。

### 熟悉golang项目目录结构

要想让你的程序跑起来，要按照这样的目录结构，正常情况下有三个目录：
```other
|--bin
|--pkg
|--src
```
其中，bin存放编译后的可执行文件；pkg存放编译后的包文件；src存放项目源文件。一般，bin和pkg目录可以不创建，go命令会自动创建（爽否？），只需要创建src目录放代码即可。

我创建一个`src`目录，下面再创建一个叫`main`的项目（可以叫任何名字，我只是示例叫`main`），里面只有一个`main.go`文件。

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-03-08-162301.jpg)

他的内容是：

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```

这样一个简单的项目就创建好了，创建好只是第一步，下面让她跑起来。

### 让她跑起来

找到配置，`Goland`里面大多数的配置都在这里。

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-03-08-164550.png)

配置你的`GOROOT`，配置成你安装的`go`路径，`Goland`会自动识别，这就是`GOROOT`的作用，和`JAVA_HOME`的作用差不多。

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-03-08-164659.jpg)

配置`GOPATH`，你的项目放在`src`下面不是随随便便就放的，得让go知道你这些个项目基于哪个位置。

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-03-08-165051.jpg)

细心的人注意到，这里有一个`Project GOPATH`，还有一个`Global GOPATH`,把你的项目配置在`Project GOPATH`里，每个项目都不一样，创建另一个项目时这个路径要配置成新项目的。

`Global GOPATH`可以弄一个公共项目，以后就把第三方的包直接装到这里，就可以**自动**在你的项目里引用了。

调出`ToolBar`，开始配置运行文件

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-03-08-162705.png)

在`ToolBar`中`Add Configuration`

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-03-09-042901.jpg)

创建一个`go build`，可以看到有一个`go remote`的选项，它是用来调试远程服务器上的代码的，有兴趣关注我，我后续更新。

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-03-09-043001.png)

注意这三个位置，

选`File`，运行文件就选`main`函数所在在文件`main.go`，输出文件夹就在和`src`同级目录的`bin`文件夹（自动创建）,`Working directory`目录就是刚刚设置`GOPATH`的目录（自动）

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-03-09-043558.jpg)

注意，如果你多次打开目录选择，框框里的目录不会被替换掉，而是追加，导致运行的时候报错，除非你想一次性编译多个项目。

例如这样：
```other
/Users/pzqu/Documents/code/go/what_go/src/main/main.go|/Users/pzqu/Documents/code/go/what_go/src/main/main.go
```

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-03-09-045127.jpg)

点击OK保存，之后，在`ToolBar`上点击运行，旁边那个符号是`debug`

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-03-09-043823.png)

成功运行！自动创建了`bin`目录

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-03-09-043858.jpg)

如果你想改输出的二进制文件名，可以在这里添加参数`-o bin/main`

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-03-09-044052.jpg)

### 如何在一个项目中使用其他项目？

####  引用自己的项目中的其他模块包

写一个新函数`func Add(a, b int) int`,放在`src`下面`main`项目，`calc`文件夹，`add.go`文件里
```tree
|____src
| |____main
| | |____calc
| | | |____add.go
| | |____main.go
```
代码如下
```go
package calc

func Add(a, b int) int {
	return  a + b
}
```

在`main`函数中调用他

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-03-09-050033.jpg)

输出结果：

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-03-09-050601.jpg)

几个点需要注意：

1. `add.go`中的Add函数名首字母必须大写， 比如Add, Addxxx.只有大写的才是Public权限，外面的包才能访问，否则只能自己文件夹下代码才能访问

 2. `add.go`的改名为addyyy.go也可以，查找add包的时候，并不会根据add.go这个文件名来查找。而是根据文件夹名来查找，一个文件夹下的所有文件都属于同一个包。所以函数变量自然不能重复。

3. `main`中调用add.Add(1,2)时，add是包， 必须跟`add.go`中的`package`处的包名一致，否则报错。

4. import后， 怎么去查找对应的包呢？ 思考一下， 很简单，无非就是GOROOT和GOPATH. 也应该明白了， src这个目录名可不是能随便取的。

#### 引用第三方项目
自己写的其他项目引入，比如我这有一个叫`common`的公共包，你的公司有可能把很多`go`包下载下来，做一个公共仓库，方便公司内网隔离。

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-03-09-144637.jpg)

代码很简单

```go
package dance

import "fmt"

func WhoDance() {
	fmt.Println("you")
}
```

在`main`里面调用
```go
package main

import "common/dance"

func main() {
	dance.WhoDance()
}
```
输出
```go
you

Process finished with exit code 0
```

还有一个相当好用的引用第三方项目的工具，`vendor`关注我的博客，我们后续再见。

### 参考

[小议并实战go包------顺便说说go中的GOROOT,GOPATH和src,pkg,bin](https://blog.csdn.net/stpeace/article/details/82710969)


