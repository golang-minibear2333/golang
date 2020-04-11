# how_to_code (go、python学习教程)

如果对这个项目感兴趣，请点击一下 **Star** :star2: :bow: :star2:， 项目会 **持续更新**，感谢大家的支持:kissing_heart:。

记录小熊工作中一时想不起:thought_balloon:的语法，或者不加班的时候做的一些小练习，:metal:用来备忘，也希望能给学习`golang`、`python`的你一些帮助:revolving_hearts:。

如果后续我用到其他语言，也有可能更新在这里。

* 哪有那么多人生开挂，不过都是厚积薄发:laughing:。
* 我的博客：https://blog.csdn.net/BTnode
* 我的个人网站: https://coding3min.com
* 我用`golang`和`python`做 [算法的项目](https://github.com/pzqu/LeetCode)
* 分享开发、运维、云技术等方面知识，编程三分钟公众号不加班的时候推送。欢迎您的关注~！

<div align="center"><img border="0" src="qrcode.jpg" alt="Coder" title="gongzhonghao" with="200" height="200"></div>

### 开始

* 如果你不会运行 golang 项目见：[让你的Golang项目在IDE里跑起来](https://coding3min.com/646.html)
* 如果你想提高编程速度见：[Goland 快捷键](goland.md)
* 如果你觉得想点击链接跳出新标签页：
    * windows: `Ctrl`+鼠标左键，在新窗口打开
    * mac: `Command` + 鼠标点击，在新窗口打开 
* 如果你想参与贡献这个项目，可以参考：[如何给开源项目贡献代码](howToContribute.md)

### 语法与简单使用

博客语法备忘快速查询：[Golang](https://coding3min.com/561.html)

我是按从易到难的顺序来排序的，有可能后面的知识会用到前面的。考虑到看的人可能想直接运行某个文件中的代码，我这里就每个都是`main`包，每个文件独立可以跑

| 单元 |                  Title                   |                  Golang                  |     Python3                |           Python2 |
| ---- | ---- | :--------------------------------------: | :--------------------------------------: |  :--------------------------------------: | 
| 变量 | 声明【变量】的各种方式 | [Golang](golang/easy/variable/variable.go) |-|-|
| |声明【常量】的各种方式 | [Golang](golang/easy/variable/const.go) |-|-|
| |类型转换|[Golang](golang/easy/type/type1.go)|-|-|
| 条件语句| switch和type switch | [Golang](golang/easy/ifelse_switch/switch.go) |-|-|
| 循环语句 | 循环语句的多种形式、死循环、break/continue | [Golang](golang/easy/for_range/for.go) |-|-| 
| range | range(范围) | [Golang](golang/easy/range/range1.go) | - | - |
| 函数|函数的简单使用| [Golang](golang/easy/function/main.go)| - | -|
| | 值传递和引用传递| [Golang](golang/easy/function/more.go)|-|-|
| | 函数当作变量使用，当做 参数传递|[Golang](golang/easy/function/function_value.go) / [实用模拟迭代器](golang/easy/function/function_value_good_demo.go)|-|-| 
| | 匿名函数与闭包| [Golang](golang/easy/function/close_package.go)| / | / |
| | 函数方法(go中定义一个类)|[Golang](golang/easy/function/go_class.go)| - | - |
| | 递归 | [Golang](golang/easy/function/recursive.go) |-|-|
| 数组| 数组定义赋值与遍历|[Golang](golang/easy/arrray/array1.go)| - | - |
|  | 多维数组 | [Golang](golang/easy/arrray/array2.go) | - | - |
| | 向函数传递数组，引用传递还是值传递？| [Golang](golang/easy/arrray/array3.go)| - | - | 
| 切片| 切片声明赋值与截取 | [Golang](golang/easy/slice/slice1.go)|-|-|
| | 切片的长度与容量，len cap append copy |  [Golang](golang/easy/slice/slice2.go)|-|-|
| 集合 | map | [Golang](golang/easy/map/map1.go)|-|-|
| 指针| 了解指针|[Golang](golang/easy/point/point1.go)| - | - |
| | 多维指针、指针作数组参数| [Golang](golang/easy/point/point2.go)| - | - |
| 结构体| 结构体 | [Golang](golang/easy/struct/struct1.go) | -|-|
| 接口 | interface 类型（接口）初识| [Golang](golang/medium/interface/interface1.go)|-|-|
| | 不实现所有方法会不会报错？| [Golang](golang/medium/interface/interface2.go)|-|-|
| | 多态| [Golang](golang/medium/interface_more/interface3.go)|-|-|
| 小工具 |代码运行时间（测速）| [Golang](golang/utils/speed.go) |-|-|
| 字符串 |快速拼接字符串|[Golang](golang/easy/string/append_string.go) / [博客应用](https://coding3min.com/675.html)|-| -|
| 错误处理 | 错误处理初识 | [Golang](golang/easy/errors/error1.go)|-|-|
|json |json解析与编码   | [Golang](golang/easy/json/parse_json.go) |-| -|
| |灵活解析多版本字段类型稍微不同的json   | [Golang](golang/medium/json_interface/fixed_json.go) | / |  / |
| 网络 | 解析域名得到ip |-| [Python3](python/network/pase_hostname.py) | 同Python3 |
| |解析url或者其中的域名 |-|- |  [Python2](python/network/py2_parse_url_hostname.py)  |
| | http 使用代理访问|[Golang](golang/medium/http_proxy/static_proxy.go)| - | - |

### `go`的并发特性

|                  Title                   |                  Golang                  |    博客讲解|
| ---- | :--------------------------------------: | :--------------------------------------: | 
| 协程（goroutine）| [Golang](golang/medium/chan/goroutine.go)| - |
| 通道（chan）的发送与接收 | [Golang](golang/medium/chan/chan.go) | - |
| select语法 | [Golang](golang/medium/chan/select.go) | - |


