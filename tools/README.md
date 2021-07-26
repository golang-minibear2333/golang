
# golang小工具

|作用|位置|备注|
|:---|:---|:---|
|代码测速|[speed.go](speed.go)||
|可超时缓存|[cache.go](cache.go)|任意类型变量缓存，可设置超时时间|
|线程安全的缓存|[map.go](map.go)|任意类型变量缓存，线程安全|
|配置proxy|[static_proxy.go](static_proxy.go)|简单配置http proxy|
|捕获panic|[defer_panic.go](defer_panic.go)||
|重试函数|[func.go](func.go)|可传入函数和重试次数，自动重试，要求返回是否成功|
| 参数|go命令行读取参数|-|-|
| 文件写入读取 |-|-|-|
|配置文件|-|-|-|

