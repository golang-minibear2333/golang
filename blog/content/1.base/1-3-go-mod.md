# 1.3 go mod最佳实践

> 本节源码位置 https://github.com/golang-minibear2333/golang/blob/master/1.base/1.3-go-mod

`java` 里有一个叫 `maven` 的包管理工具， `go` 也有一个叫 `go mod` 的管理工具，可以管理项目引用的第三方包版本、自动识别项目中用到的包、自动下载和管理包。

为什么要使用go mod？

* 使用go mod仓库中可以不用再上传依赖代码包，防止代码仓库过大浪费以及多个项目同时用包时的浪费
* 可以管理引用包的版本，这一点是gopath（src模式）和vendor做不到的
* 如果依赖gopath不同项目如果引用了同一个软件包的不同版本，就会造成编译麻烦

gopath是go之前的默认策略，每个项目在运行时都要严格放在`src`目录下，而go mod不用

**原来的包管理方式**

- 在不使用额外的工具的情况下，`Go` 的依赖包需要手工下载，
- 第三方包没有版本的概念，如果第三方包的作者做了不兼容升级，会让开发者很难受
- 协作开发时，需要统一各个开发成员本地`$GOPATH/src`下的依赖包
- 引用的包引用了已经转移的包，而作者没改的话，需要自己修改引用。
- 第三方包和自己的包的源码都在`src`下，很混乱。对于混合技术栈的项目来说，目录的存放会有一些问题

**新的包管理模式解决了以上问题**

- 自动下载依赖包
- 项目不必放在`$GOPATH/src`内了
- 项目内会生成一个`go.mod`文件，列出包依赖
- 所以来的第三方包会准确的指定版本号
- 对于已经转移的包，可以用 `replace` 申明替换，不需要改代码

## 1.3.1 配置

`golang>=1.12`
添加环境变量 `GO111MODULE` 为 `on` 或者 `auto` ，设置方法

```BASH
go env GO111MODULE=on
```

```shell
go env -w GO111MODULE="on"
go env -w GOPROXY=https://goproxy.io
go mod init 项目名
go mod tidy #add missing and remove unused modules
```

* 打开go mod 模式
* 使用国内下载包代理
* 初始化mod项目
* 自动增加包和删除无用包到 GOPATH 目录下（build的时候也会自动下载包加入到go.mod里面的）

注意：只要在本地设置一个公用path目录就可以了，全部的包都会下载到那里，其他本地项目用到时就可以共享了

自动生成了go.mod和go.sum文件，可以不用理会，下面是简单介绍

## 1.3.2 go.mod 文件

go.mod 的内容比较容易理解

* 第一行：模块的引用路径
* 第二行：项目使用的 go 版本
* 第三行：项目所需的直接依赖包及其版本

在实际应用上，你会看见更复杂的 go.mod 文件，比如下面这样

```go
module github.com/BingmingWong/module-test

go 1.14

require (
example.com/apple v0.1.2
example.com/banana v1.2.3
example.com/banana/v2 v2.3.4
example.com/pear // indirect
example.com/strawberry // incompatible
)

exclude example.com/banana v1.2.4
replace（
golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac = > github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
golang.org/x/net v0.0.0-20180821023952-922f4815f713 = > github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
golang.org/x/text v0.3.0 = > github.com/golang/text v0.3.0
)
```

主要是多出了两个 flag：

* exclude：忽略指定版本的依赖包
* replace：由于在国内访问golang.org/x的各个包都需要翻墙，你可以在go.mod中使用replace替换成github上对应的库。

## 1.3.3 go.sum 文件

每一行都是由 模块路径，模块版本，哈希检验值 组成，其中哈希检验值是用来保证当前缓存的模块不会被篡改。hash 是以h1:开头的字符串，表示生成checksum的算法是第一版的hash算法（sha256）。

值得注意的是，为什么有的包只有一行

```shell
<module> <version>/go.mod <hash>
```

而有的包却有两行呢

```shell
<module> <version> <hash>
<module> <version>/go.mod <hash>
```

那些有两行的包，区别就在于 hash 值有两行，一行是 `h1:hash` 也就是模块包的hash，另一行是 go.mod `h1:hash`，举例如下

```shell
github.com/sirupsen/logrus v1.8.1 h1:dJKuHgqk1NNQlqoA6BTlM1Wf9DOH3NBjQyu0h9+AZZE=
github.com/sirupsen/logrus v1.8.1/go.mod h1:yWOB1SBYBC5VeMP7gHvWumXLIWorT60ONWic61uBYv0=
```

而  `h1:hash` 和 `go.mod` `h1:hash`两者，要不就是同时存在，要不就是只存在 `go.mod` `h1:hash`。那什么情况下会不存在 `h1:hash` 呢，就是当 Go
认为肯定用不到某个模块版本的时候就会省略它的`h1 hash`，就会出现不存在 `h1 hash`，只存在 `go.mod` `h1:hash` 的情况。

`go.mod` 和 `go.sum` 是 go modules 版本管理的指导性文件，因此 `go.mod` 和 `go.sum` 文件都应该提交到你的 Git 仓库中去，避免其他人使用你写项目时，重新生成的`go.mod`
和 `go.sum` 与你开发的基准版本的不一致。

## 1.3.4 go mod 命令的使用

go mod init：初始化go mod， 生成go.mod文件，后可接参数指定 module 名，上面已经演示过。

go mod download：手动触发下载依赖包到本地cache（默认为$GOPATH/pkg/mod目录）

go mod graph：打印项目的模块依赖结构

go mod tidy ：添加缺少的包，且删除无用的包

go mod verify ：校验模块是否被篡改过

go mod why：查看为什么需要依赖

go mod vendor ：导出项目所有依赖到vendor下

写入go.mod有两种方法：

* 你只要在项目中有 import 并使用或者使用下划线强制占用，然后 go build 就会 go module 就会自动下载并添加。
* `go mod tidy`

## 1.3.5 vendor是什么

vendor是项目缓存，为了防止开源代码项目被删除无法引用下载，会使用vendor来做缓存管理，它是独立的，你可以手动管理引用的包，代码包查找的顺序是向上冒泡

```shell
包同目录下的vendor
包目录向上的最近的一个vendor
...
GOPATH src 下的vendor
GOROOT src
GOPATH src
```

这样的话， 我们可以把包的依赖都放在 vendor 下，然后提交到仓库，这样可以省却拉取包的时间，并且相对自由，你想怎么改都可以

## 1.3.6 最佳实践

go mod 只是一个依赖包版本管理工具，包的查找顺序还是一样的，使用mod就不用把代码都放到src下来管理，可以根据go.mod文件中记录的版本来索引

我建议：

* 使用mod管理版本，并使用go vendor来cache依赖包，上传到仓库防止代码包被删除
* 运行时用到自己的项目，不要使用本地代码，而是保证依赖包都是稳定的，防止忘记提交
* 如果你想发布包把自己写的模板给别人用，记得提交到仓库

这样就可以单个项目独立下来debug了，依赖包版本也管理上了

PS: go项目就可以使用mod和vendor，如果要集成其他语言代码为子模块可以使用git submodule

## 1.3.7 tips

**Q1: 我的包下哪去了？**

A: 依赖的第三方包被下载到了 `$GOPATH/pkg/mod` 路径下。

**Q2: `GO111MODULE` 的三个参数 `auto` 、 `on` 、 `off` 有什么区别？**

A: `auto` 根据是否在 `src` 下自动判定， `on` 只用 `go.mod` ， `off` 只用 `src` 。

**Q3: 依赖包中的地址失效了怎么办？比如 golang. org/x/… 下的包都无法下载怎么办？**

A: 在 `go.mod` 文件里用 `replace` 替换包，例如

```
replace golang.org/x/text => github.com/golang/text latest
```

这样， `go` 会用 `github.com/golang/text` 替代 `golang.org/x/text`

**Q4: 在 `go mod` 模式中，项目自己引用自己中的某些模块怎么办？**

A: `go.mod` 文件里的第一行会申明 `module main` ，把这个 `main` 改为你的项目名，引用的时候就 `import "项目名/模块名"` 即可。

> 根据官方的说法，从 `Go 1.13` 开始，模块管理模式将是 Go 语言开发的**默认模式**。


## 1.3.8 小结

go mod 是未来的默认模式，未来会取消 go path 也就是src 的方式，但自己的项目目录还是尽量按路径放置，不然回头找不到了

## 1.3.9 引用

[Go Modules](https://mp.weixin.qq.com/s/U2teH_ZCSMW6qNH404t6yw)
[掘金](https://juejin.im/post/5c9c8c4fe51d450bc9547ba1)
