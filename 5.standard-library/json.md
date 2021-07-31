# json处理

日常工作中，最常用的数据传输格式就是`json`，而`encoding/json`库是内置做解析的库。这一节来看看它的用法，还有几个日常使用中隐晦的陷阱和处理技巧。

## json与struct

一个常见的接口返回内容如下：

```json
{
  "data": {
    "items": [
      {
        "_id": 2
      }
    ],
    "total_count": 1
  },
  "message": "",
  "result_code": 200
}
```

在`golang`中往往是要把`json`格式转换成结构体对象使用的。 新版`Goland`粘贴`json`会自动生成结构体，也可以在 http://json2struct.mervine.net 获得 `json` 到`struct`的自动转换。

```Go
type ResponseData struct {
	Data struct {
		Items []struct {
			Id int `json:"_id"`
		} `json:"items"`
		TotalCount int `json:"total_count"`
	} `json:"data"`
	Message    string `json:"message"`
	ResultCode int    `json:"result_code"`
}
```

用反斜杠加注解的方式表明属于`json`中哪个字段，要注意不应该嵌套层数过多，难以阅读容易出错。

一般把内部结构体提出来，可能会另做他用。

```go
type ResponseData struct {
	Data struct {
		Items []Body `json:"items"`
		TotalCount int64 `json:"total_count"`
	} `json:"data"`
	Message    string `json:"message"`
	ResultCode int64  `json:"result_code"`
}

type Body struct {
	ID int `json:"_id"`
}
```

## 解析

解析就是`json`字符串，转`struct`类型。如下，第一个参数为字节数组，第二个为接收的结构体实体地址。如有报错返回错误信息，如没有返回`nil`。

```go
//函数签名
func Unmarshal(data []byte, v interface{}) error
// 用法
err := json.Unmarshal([]byte(jsonStr), &responseData)
```

完整代码如下

```go
func foo() {
	jsonStr := `{"data":{"items":[{"_id":2}],"total_count":1},"message":"","result_code":200}`
	//把string解析成struct
	var responseData ResponseData
	err := json.Unmarshal([]byte(jsonStr), &responseData)
	if err != nil {
		fmt.Println("parseJson error:" + err.Error())
		return
	}
	fmt.Println(responseData)
}
```

输出如下，和`java`的`toString`不同，直接输出了值，如有需要要自行实现并绑定`ToString`方法。

```go
{{[{2}] 1}  200}
```


## 反解析

第一步，复习初始化结构体的方法。

```go
r := ResponseData{
    Data: struct {
        Items      []Body `json:"items"`
        TotalCount int64  `json:"total_count"`
    }{
        Items: []Body{
            {ID: 1},
            {ID: 2},
        },
        TotalCount: 1,
    },
    Message:    "",
    ResultCode: 200,
}
```

如上，无类型的结构体`Data`需要明确把类型再写一遍，再为其赋值。`[]Body`内直接赋值列表。

反解析函数签名如下，传入结构体，返回编码好的`[]byte`，和可能的报错信息。

```go
func Marshal(v interface{}) ([]byte, error)
```

完整代码如下

```go
func bar() {
	r := ResponseData{
		....
	}
	//把struct编译成string
	resBytes, err := json.Marshal(r)
	if err != nil {
		fmt.Println("convertJson error: " + err.Error())
	}
	fmt.Println(string(resBytes))
}
```

输出

```go
{"data":{"items":[{"_id":1},{"_id":2}],"total_count":1},"message":"","result_code":200}
```

## 陷阱1、忘记取地址

解析的代码，结尾处应该是`&responseData)` 忘记取地址会导致无法赋值成功，返回报错。

```go
err := json.Unmarshal([]byte(jsonStr), responseData)
```

输出报错

```go
json: Unmarshal(non-pointer main.ResponseData)
```

## 陷阱2、大小写

定义一个简单的结构体来演示这个陷阱。

```go
type People struct {
	Name string `json:"name"`
	age  int    `json:"age"`
}
```

变量如果需要被外部使用，也就是`java`中的`public`权限，定义时首字母用大写就可以实现。

```go
type People struct
```

要用来解析`json`的`struct`内部假如使用了小写作为变量名，会导致无法解析成功，而且不会报错！

```go
func err1() {
	reqJson := `{"name":"minibear2333","age":26}`
	var person People
	err := json.Unmarshal([]byte(reqJson), &person)
	if err != nil {...}
	fmt.Println(person)
}
```

输出，没有成功取到`age`字段。

```go
{minibear2333 0}
```

这是因为标准库中是使用反射来获取的，私有字段是无法获取到的，所以内部不知道有这个字段，自然无法显示报错信息。

我以前没有用自动解析，手敲上去结构体，很容易出现这样的问题，无论如何漏掉某个字段。好在编译器会有提示。

![](https://coding3min.oss-accelerate.aliyuncs.com/2021/07/31/R5vaRr.png)


## 陷阱3、十六进制或其他非 UTF8 字符串

Go 默认使用的字符串编码是 UTF8 编码的。直接解析会出错

```go
func err2() {
	raw := []byte(`{"name":"\xc2"}`)
	var person People
	if err := json.Unmarshal(raw, &person); err != nil {
		fmt.Println(err)
	}
}
```

输出
```go
invalid character 'x' in string escape code
```

加上反斜杠转义可以成功，或者使用`base64`编码成字符串，要特别注意这一点，一下子就体现出单元测试的重要性了。

```go
raw := []byte(`{"name":"\\xc2"}`)
raw := []byte(`{"name":"wg=="}`)
```

其他需要注意的是编码如果不是`UTF-8`格式，那么`Go`会用 `�` (`U+FFFD`) 来代替无效的 UTF8，不会报错，但是获得的字符串可能不是你需要的结果。

## 陷阱四、数字转interface{}

因为默认编码无类型数字视为 `float64` 。如果想用类型判断语句为`int`会直接`panic`。

```go
func err4() {
	var data = []byte(`{"age": 26}`)
	var result map[string]interface{}
	...
	var status = result["age"].(int) //error
}
```

* 上面的代码隐含一个知识点，`json`中`value`是简单类型时，可以直接解析成字典。
* 如果有嵌套，那么内部类型也会解析成字典。

运行时 Panic:

```go
panic: interface conversion: interface {} is float64, not int

goroutine 1 [running]:
main.err4()
```

* 可以先转换成`float64`再转换成`int`
* 其实还有几种方法，太麻烦了没有必要，就不做介绍了。

## 技巧、版本变更兼容

你有没有遇到过一种场景，一个接口更新了版本，把`json`的某个字段变更了，在请求的时候每次都定义两套`struct`。

比如`Age`在版本1中是`int`在版本2中是`string`，解析的过程中就会出错。

```go
json: cannot unmarshal number into Go struct field People.age of type string
```

我在下面介绍一个技巧，可以省去每次解析都要转换的工作。

我在源码里面看到，无论反射获得的是哪种类型都会去调用相应的解析接口`UnmarshalJSON`。

结合前面的知识，在`Go`里面看起来像鸭子就是鸭子，我们只要实现这个方法，并绑定到结构体对象上，就可以让源码来调用我们的方法。

```go
type People struct {
    Name string `json:"name"`
    Age  int    `json:"_"`
}
func (p *People) UnmarshalJSON(b []byte) error {
	...
}
```

* 使用下划线，表示此类型不解析。
* 必须用指针的方式绑定方法。
* 必须与interface{}中定义的方法签名完全一致。

一共有四个步骤

1、定义临时类型。用来接受非`json:"_"`的字段，注意用的是`type`关键字。

```go
type tmp People
```

2、用中间变量接收json串，tmp以外的字段用来接受`json:"_"`属性字段

```go
var s = &struct {
    tmp
    // interface{}类型，这样才可以接收任意字段
    Age interface{} `json:"age"`
}{}
// 解析
err := json.Unmarshal(b, &s)
```

3、判断真实类型，并类型转换

```go
switch t := s.Age.(type) {
case string:
    var age int
    age, err = strconv.Atoi(t)
    if err != nil {...}
    s.tmp.Age = age
case float64:
    s.tmp.Age = int(t)
}
```

4、tmp类型转换回People，并赋值
```go
*p = People(s.tmp)
```

## 小结

通过本节，我们掌握了标准库中`json`解析和反解析的方法，以及很有可能日常工作中踩到的几个坑。

最后分享了一个技巧，实际使用中，这个技巧更加灵活。

留一个作业：假如有`v1`和`v2`不同的两个版本`json`几乎完成不同，业务逻辑已经使用`v1`版本，是否可以把`v2`版本转换成`v1`版本，不用改动业务逻辑？

提示： 可以通过深拷贝把`v2`版本解析出来的结构体完全转换成`v1`版本的结构体。
要求：必须使用实现 `UnmarshalJSON`的技巧。
