# 1.2.1 hello world

Go 语言是谷歌 2009 发布的第二款开源编程语言。  
Go 语言专门针对多处理器系统应用程序的编程进行了优化，使用 Go 编译的程序可以媲美 C 或 C++代码的速度，而且更加安全、支持并行进程。

## 1.2.1 为什么要选择学习 Go 语言呢？与其他语言的应用相比，它有什么优点呢？

1、学习曲线它包含了类 C 语法、GC 内置和工程工具。这一点非常重要，因为 Go 语言容易学习，所以一个普通的大学生花一个星期就能写出来可以上手的、高性能的应用。在国内大家都追求快，这也是为什么国内 Go 流行的原因之一。

2、效率 Go 拥有接近 C 的运行效率和接近 PHP 的开发效率，这就很有利的支撑了上面大家追求快速的需求。

3、出身名门、血统纯正之所以说 Go 语言出身名门，是因为我们知道 Go 语言出自 Google 公司，这个公司在业界的知名度和实力自然不用多说。Google 公司聚集了一批牛人，在各种编程语言称雄争霸的局面下推出新的编程语言，自然有它的战略考虑。而且从 Go 语言的发展态势来看，Google 对它这个新的宠儿还是很看重的，Go 自然有一个良好的发展前途。我们看看 Go 语言的主要创造者，血统纯正这点就可见端倪了。

4、自由高效：组合的思想、无侵入式的接口 Go 语言可以说是开发效率和运行效率二者的完美融合，天生的并发编程支持。Go 语言支持当前所有的编程范式，包括过程式编程、面向对象编程以及函数式编程。程序员们可以各取所需、自由组合、想怎么玩就怎么玩。

5、强大的标准库这包括互联网应用、系统编程和网络编程。Go 里面的标准库基本上已经是非常稳定了，特别是我这里提到的三个，网络层、系统层的库非常实用。

6、部署方便：二进制文件、Copy 部署我相信这一点是很多人选择 Go 的最大理由，因为部署太方便了，所以现在也有很多人用 Go 开发运维程序。

7、简单的并发它包含了降低心智的并发和简易的数据同步，我觉得这是 Go 最大的特色。之所以写正确的并发、容错和可扩展的程序如此之难，是因为我们用了错误的工具和错误的抽象，Go 可以说这一块做的相当简单。

8、稳定性 Go 拥有强大的编译检查、严格的编码规范和完整的软件生命周期工具，具有很强的稳定性，稳定压倒一切。那么为什么 Go 相比于其他程序会更稳定呢？这是因为 Go 提供了软件生命周期（开发、测试、部署、维护等等）的各个环节的工具，如 go tool、gofmt、go test。

## 1.2.2 Go 语言适合用来做什么？

服务器编程：以前你如果使用 C 或者 C++做的那些事情，用 Go 来做很合适，例如处理日志、数据打包、虚拟机处理、文件系统等。

分布式系统：数据库代理器等。

网络编程：这一块目前应用最广，包括 Web 应用、API 应用、下载应用、内存数据库。

云平台：google 开发的 groupcache，couchbase 的部分组建云平台，目前国外很多云平台在采用 Go 开发，CloudFoundy 的部分组建，前 VMare 的技术总监自己出来搞的 apcera 云平台。

## 1.2.3 Go 语言成功的项目

nsq：bitly 开源的消息队列系统，性能非常高，目前他们每天处理数十亿条的消息  
docker：基于 lxc 的一个虚拟打包工具，能够实现 PAAS 平台的组建  
packer：用来生成不同平台的镜像文件，例如 VM、vbox、AWS 等，作者是 vagrant 的作者  
skynet：分布式调度框架  
Doozer：分布式同步工具，类似 ZooKeeper  
Heka：mazila 开源的日志处理系统  
cbfs：couchbase 开源的分布式文件系统  
tsuru：开源的 PAAS 平台，和 SAE 实现的功能一模一样  
groupcache：memcahe 作者写的用于 Google 下载系统的缓存系统  
god：类似 redis 的缓存系统，但是支持分布式和扩展性  
gor：网络流量抓包和重放工具

## 1.2.4 哪些大公司在用 go 语言？

Google  

这个不用多做介绍，作为开发 Go 语言的公司，当仁不让。Google 基于 Go 有很多优秀的项目，比如：https://github.com/kubernetes/kubernetes ，大家也可以在 Github 上 https://github.com/google/ 查看更多 Google 的 Go 开源项目。

Facebook  

Facebook 也在用，为此他们还专门在 Github 上建立了一个开源组织 facebookgo，大家可以通过 https://github.com/facebookgo 访问查看 facebook 开源的项目，比如著名的是平滑升级的 grace。腾讯  
腾讯作为国内的大公司，还是敢于尝试的，尤其是 Docker 容器化这一块，他们在 15 年已经做了 docker 万台规模的实践，具体可以参考 http://www.infoq.com/cn/articles/tencent-millions-scale-docker-application-practice

百度  

目前所知的百度的使用是在运维这边，是百度运维的一个 BFE 项目，负责前端流量的接入。他们的负责人在 2016 年有分享，大家可以看下这个 http://www.infoq.com/cn/presentations/application-of-golang-in-baidu-frontend

阿里  

阿里巴巴具体的项目不太清楚，不过听说其系统部门、CDN 等正在招 Go 方面的人。京东  
京东云消息推送系统、云存储，以及京东商城等都有使用 Go 做开发。

小米  

小米对 Golang 的支持，莫过于运维监控系统的开源，也就是 http://open-falcon.com/

此外，小米互娱、小米商城、小米视频、小米生态链等团队都在使用Golang。

360  

360 对 Golang 的使用也不少，一个是开源的日志搜索系统 Poseidon，托管在 Github 上，https://github.com/Qihoo360/poseidon

Go语言前景：

![](https://coding3min.oss-accelerate.aliyuncs.com/2020/06/07/6cQ1wK1550.jpg)

![](https://coding3min.oss-accelerate.aliyuncs.com/2020/06/07/BrcWXA1550.png)

（以上数据来源于网络）

## 1.2.5 第一个 go 程序

带着目标学东西往往是最有成效的，**为什么学**以及**环境安装**可以参考前面的文章。

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

## 1.2.6 让你的项目在IDE里跑起来


每次新建项目，不熟悉go的项目结构，一般跑都跑不起来，每次都要重新搞一遍，好几回跑项目都会报类似`File is  invalid`的错误

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-03-08-145230.jpg)

有时候报其他奇怪的错误，今天就下决心整理一下，理一理概念 GOROOT、GOPATH、src、 pkg、bin，希望以后不要再出现这样的问题了，同时给看到文章的你一些帮助。

### 1.2.6.1 熟悉golang项目目录结构

要想让你的程序跑起来，要按照这样的目录结构，正常情况下有三个目录：
```other
|--bin
|--pkg
|--src
```
其中，bin存放编译后的可执行文件；pkg存放编译后的包文件；src存放项目源文件。一般，bin和pkg目录可以不创建，go命令会自动创建（爽否？），只需要创建src目录放代码即可。

我创建一个`src`目录，下面再创建一个叫`main`的项目（可以叫任何名字，我只是示例叫`main`），里面只有一个`main.go`文件。

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-03-08-162301.2.jpg)

他的内容是：

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```

> 本节源码位置 https://github.com/golang-minibear2333/golang/blob/master/1.base/1.2-hello-world


这样一个简单的项目就创建好了，创建好只是第一步，下面让她跑起来。

### 1.2.6.2 让她跑起来

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

![](https://coding3min.oss-accelerate.aliyuncs.com/coding3min/2020-03-09-044051.2.jpg)

## 1.2.7 如何在一个项目中使用其他项目？

### 1.2.7.1 引用自己的项目中的其他模块包

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

### 1.2.7.2 引用第三方项目

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

## 1.2.8 小结

通过这一节，你已经了解到了go语言的历史和前景，并了解到怎么在IDE里跑起来go项目。这是一切的开始，算是进入了go语言的大门，在接下来的日子希望我们可以愉快的走下去。

## 1.2.9 参考

[小议并实战go包------顺便说说go中的GOROOT,GOPATH和src,pkg,bin](https://blog.csdn.net/stpeace/article/details/82710969)










