### 简介

`kingpin` 功能比 `flag` 库强大，用法差不多。
相比 `flag` 库，最重要的一点就是支持不加 `-` 的调用。
比如一个命令行程序有三个函数分别为 `A` , `B` , `C` ，要实现根据命令行的输入运行不同的函数，如果用flag实现的话应该是下面这种使用方法：

``` BASH
./cli --method A
./cli --method B
./cli --method C
```

每次都需要输入 `--method` ，然而用 `kingpin` 库实现的话就可以达到下面这种效果：
``` BASH 
./cli A
./cli B
./cli C

``` 
节省了很多输入操作。

### 使用方法

``` BASH
go get gopkg.in/alecthomas/kingpin.v2
go mod vendor
```

这样子 `go.mod` 文件里就引入了， `vendor` 文件夹就缓存了此包，然后直接在代码中使用。

``` Go
package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
)

func main() {
	var (
		listenAddress = kingpin.Flag(
			"web.listen-address",
			"Address on which to expose metrics and web interface.",
		).Default(":18001").String()
		metricsPath = kingpin.Flag(
			"web.telemetry-path",
			"Path under which to expose metrics.",
		).Default("/metrics").String()
		
	)
	kingpin.HelpFlag.Short('h')
    kingpin.Parse()
    
	conf.ApiMtncUrl = *apiMtncPath

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Node Exporter</title></head>
			<body>
			<h1>xxx Exporter</h1>
			<p><a href=" ` + *metricsPath + ` ">Metrics</a></p>
			</body>
			</html>`))
	})

	http.Handle("/metrics", XXXX.Handler())

	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		fmt.Printf("Error occur when start server %v", err)
	}
}
```

官方文档参考 [package kingpin
](http://godoc.org/gopkg.in/alecthomas/kingpin.v2)

### 引用

[Golang命令行参数解析库kingpin](https://xuanyu.li/2017/08/05/golang-cli-args-parse/)
