### 引言

对于dockerfile而言，何为完美? 我认为应该满足以下三点：

* 体积小
* 构建快
* 够安全

PS: 注意！从 `Docker 17.05` 版本起， `Docker` 才开始支持容器镜像的多阶段构建(multi-stage build)，所以本文所使用 `docker` 版本必须高于 `17.05` （多阶段构建的意思就是把编译的过程也放同一个 `Dockerfile` 里，不用在自己的开发机或服务器上编译，再把编译出的二进制程序打入镜像）

### 可联网的环境

> 根据官方的说法，从 `Go 1.13` 开始，模块管理模式将是 Go 语言开发的**默认模式**。

我们使用go mod 做包管理，就不需要有任何额外配置

``` Dockerfile
FROM golang:1.13.5-alpine3.10 AS builder

WORKDIR /build
RUN adduser -u 10001 -D app-runner

ENV GOPROXY https://goproxy.cn
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o your-application .

FROM alpine:3.10 AS final

WORKDIR /app
COPY --from=builder /build/your-application /app/
#COPY --from=builder /build/config /app/config
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER app-runner
ENTRYPOINT ["/app/your-application"]
```

**逐行拆解**
这里的拆解完全引用 [手把手教你写一个完美的Golang Dockerfile](https://studygolang.com/articles/26823)

首先，这个dockerfile分为builder和final两部分。

**builder** 选择了 `golang:1.13.5-alpine3.10` 作为编译的基础镜像，相比于 `golang:1.13` , 一方面是因为它体积更小，另一方面是我发现 `golang:1.13` 的编译结果，在 `alpine:3.10` 中会报 `not found` 的错误，虽说有人提供了其它的解决方案，但是能直接避免，为啥不避免呢。

``` Dockerfile
RUN adduser -u 10001 -D app-runner
``` 
接着是创建了一个 `app-runner` 的用户, `-D` 表示无密码。

此用户的信息是是需要拷到 `final` 中，作为应用程序的启动用户。这是为了避免使用 `container` 中的默认用户 `root` ，那可是有安全漏洞的，详细解释，可以参考这篇 `medium` 上的文章[Processes In Containers Should Not Run As Root](https://medium.com/@mccode/processes-in-containers-should-not-run-as-root-2feae3f0df3b)

再下面的四行，

``` Dockerfile
ENV GOPROXY https://goproxy.cn
COPY go.mod .
COPY go.sum .
RUN go mod download
```

是配置了国内的代理，安装依赖包了。这里用 `go mod download` 的好处是下次构建镜像文件时，当go.mod和go.sum没有改变时，它是有缓存的，可以避免重复下载依赖包，加快构建。

builder的最后，就是把当前目录的文件拷过去，编译代码了。

``` Dockerfile
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o your-application .
``` 

`final` 选择了 `alpine:3.10` ,一方面是体积小，只有 `5m` ；另一方面也是和构建镜像的 `alpine` 版本保持一致。

接下来几行没啥说的，就是把构建结果、配置文件（有的话）和用户的相关文件拷过去。

下面的这步一定不要忘记了，
``` Dockerfile
USER app-runner
``` 
没有它， `container` 启动时就是用 `root` 用户启动了!!! 如果被攻击了，那黑客可是就有 `root` 权限了（不要问我为啥会被攻击）。

最后，设置一个 `ENTRYPOINT` ，完事!

如果你程序的启动过程比较复杂，或者是要在启动时根据环境变量的值做不同的操作，那还是写个 `shell` 文件吧。

### 离线打包

``` Dockerfile
# Building stage
FROM golang:1.13.5-alpine3.10 AS builder

WORKDIR /build/src/your-application
RUN adduser -u 10001 -D app-runner

ENV GO111MODULE off
ENV GOPATH /build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o your-application  main.go
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o your-application main.go

# Production stage
FROM alpine:3.10 AS final

WORKDIR /app

COPY --from=builder /build/src/your-application/example/linux /app
COPY --from=builder /build/src/your-application/your-application /app
#COPY --from=builder /build/src/your-application/conf /app/conf

RUN adduser -u 10001 -D app-runner
RUN chmod -R 755 /app

ENTRYPOINT ["/app/your-application"]
```

如果你的环境是内网，不能连接外网（不能联网），要从外部导入一个 `go mod` 项目，并运行的时候，肯定会 `timeout` 在下载项目依赖的包的阶段，实际上依赖包已经放到目录文件，不用下载也能正常运行。为了解决这一问题，我们只需要设置参数 `GO111MODULE=off` ，然后设置正确的 `GOPATH` 即可

``` Dockerfile
ENV GO111MODULE off
ENV GOPATH /build
```

在代码库中需要提前把代码包的中 `vendor` 更新，在本地执行以下命令，并提交到代码库

``` BASH 
go mod init your-application
go mod vendor

``` 

这样就会有离线的 `vendor` 代码库

```

|——vendor

    └──github.com
    └──golang.org
    └──gopkg.in
    └──modules.txt

``` 

* `GO111MODULE=off` 无 `mod` 支持， `go` 会从 `GOPATH` 和 `vendor` 文件夹寻找包。
* `GO111MODULE=on` 模块支持，go 会忽略 `GOPATH` 和 `vendor` 文件夹，只根据 `go.mod` 下载依赖。
* `GO111MODULE=auto` 在 `$GOPATH/src` 外面且根目录有 `go.mod` 文件时，开启模块支持。

### 有可能会遇到的问题

#### docker镜像源速度慢

如果docker镜像拉取速度太慢，或者拉取不到，可以试试改为国内镜像源地址，参考[这里](https://coding3min.com/1229.html)

#### 更新docker的yum源

如果你发现自己的`docker`版本低，但是自己的源里面又没有想要的版本，那就需要更新官方的源
参考[这里](https://coding3min.com/1227.html)

### 引用 

[手把手教你写一个完美的Golang Dockerfile](https://studygolang.com/articles/26823)
[Golang1.5到Golang1.12包管理：golang vendor 到 go mod](https://studygolang.com/articles/18670)
[官方golang包管理神器，值得一试！go mod | 编程三分钟](https://coding3min.com/801.html)


