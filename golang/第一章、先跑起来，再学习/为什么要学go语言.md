> go语言github项目：https://github.com/minibear2333/how_to_code

![](https://coding3min.oss-accelerate.aliyuncs.com/2020/06/07/g7bbHb1550.jpg)

Go 语言是谷歌 2009 发布的第二款开源编程语言。  
Go 语言专门针对多处理器系统应用程序的编程进行了优化，使用 Go 编译的程序可以媲美 C 或 C++代码的速度，而且更加安全、支持并行进程。

[TOC]

### 为什么要选择学习 Go 语言呢？与其他语言的应用相比，它有什么优点呢？

1、学习曲线它包含了类 C 语法、GC 内置和工程工具。这一点非常重要，因为 Go 语言容易学习，所以一个普通的大学生花一个星期就能写出来可以上手的、高性能的应用。在国内大家都追求快，这也是为什么国内 Go 流行的原因之一。

2、效率 Go 拥有接近 C 的运行效率和接近 PHP 的开发效率，这就很有利的支撑了上面大家追求快速的需求。

3、出身名门、血统纯正之所以说 Go 语言出身名门，是因为我们知道 Go 语言出自 Google 公司，这个公司在业界的知名度和实力自然不用多说。Google 公司聚集了一批牛人，在各种编程语言称雄争霸的局面下推出新的编程语言，自然有它的战略考虑。而且从 Go 语言的发展态势来看，Google 对它这个新的宠儿还是很看重的，Go 自然有一个良好的发展前途。我们看看 Go 语言的主要创造者，血统纯正这点就可见端倪了。

4、自由高效：组合的思想、无侵入式的接口 Go 语言可以说是开发效率和运行效率二者的完美融合，天生的并发编程支持。Go 语言支持当前所有的编程范式，包括过程式编程、面向对象编程以及函数式编程。程序员们可以各取所需、自由组合、想怎么玩就怎么玩。

5、强大的标准库这包括互联网应用、系统编程和网络编程。Go 里面的标准库基本上已经是非常稳定了，特别是我这里提到的三个，网络层、系统层的库非常实用。

6、部署方便：二进制文件、Copy 部署我相信这一点是很多人选择 Go 的最大理由，因为部署太方便了，所以现在也有很多人用 Go 开发运维程序。

7、简单的并发它包含了降低心智的并发和简易的数据同步，我觉得这是 Go 最大的特色。之所以写正确的并发、容错和可扩展的程序如此之难，是因为我们用了错误的工具和错误的抽象，Go 可以说这一块做的相当简单。

8、稳定性 Go 拥有强大的编译检查、严格的编码规范和完整的软件生命周期工具，具有很强的稳定性，稳定压倒一切。那么为什么 Go 相比于其他程序会更稳定呢？这是因为 Go 提供了软件生命周期（开发、测试、部署、维护等等）的各个环节的工具，如 go tool、gofmt、go test。

### Go 语言适合用来做什么？

服务器编程：以前你如果使用 C 或者 C++做的那些事情，用 Go 来做很合适，例如处理日志、数据打包、虚拟机处理、文件系统等。

分布式系统：数据库代理器等。

网络编程：这一块目前应用最广，包括 Web 应用、API 应用、下载应用、内存数据库。

云平台：google 开发的 groupcache，couchbase 的部分组建云平台，目前国外很多云平台在采用 Go 开发，CloudFoundy 的部分组建，前 VMare 的技术总监自己出来搞的 apcera 云平台。

### Go 语言成功的项目

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

### 哪些大公司在用 go 语言？

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







