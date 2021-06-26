# 3.2 结构体

> 本节源码位置 https://github.com/golang-minibear2333/golang/blob/master/3.grammar-advancement/3.2-struct

## 3.2.1 go 语言中的结构体

和 `c++` 的结构体类似，如下定义一个结构体类型。

```Go
type Body struct {
	name string
	age  int
}
```

像这样就可以使用

```Go
var body Body
body.name = "coding3min"
body.age = 12
fmt.Println(body)
```

输出

```
{coding3min 12}
```

## 3.2.2 go 中的类

结构体在 `go` 中是最常用的一种语法，有没有想过为什么？

这是因为我们学过一些面向对象的语言，其中有一个叫类的概念，但是 `go` 里面没有。

`go` 用一种特殊的方式，把结构体本身看作一个类。

一个成熟的类，具备成员变量和成员函数，结构体本身就有成员变量，再给他绑定上成员函数，是不是就可以了！

```Go
type people struct {
	name string
}

func (p people) toString() {
	fmt.Println(p.name)
	fmt.Printf("p的地址 %p \n", &p)
}
```

上面给 `people` 结构体绑定了一个函数, 调用下看看

```Go
p1 := people{"coding3min"}
p1.toString()
```

按照 `toString()` 方法的内容，先输出 `name` 再输出 `p的地址`

```Go
coding3min
p的地址 0xc0001021f0  #这里的地址一会有用
```

再绑定一个函数，你想想和上面的函数有什么区别，注意 `60%` 的人第一眼都没看出来

```Go
func (p *people) sayHello() {
	fmt.Printf("Hello! %v \n", p.name)
	fmt.Printf("*p的地址 %p \n", p)
}
```

可以注意到，和 `toString()` 函数不同的是， `sayHello()` 用了指针的方式进行绑定。

输出，可以注意到这里的地址和上面的不同。

```Go
Hello! coding3min
*p的地址 0xc00008e1e0
```

这两种绑定方式，都是相当于给结构体绑定了函数，这个结构体等价于对象，唯一的不同点就是如果使用 `*` 绑定函数，那么这种对象就是单例的，引用的是同一个结构体。

```Go
p1 := people{"coding3min"}
p1.sayHello()
p2 := &people{"tom"}
p2.sayHello()
```

输出，可以看到地址一致。

```Go
*p的地址 0xc00008e220
p2的地址 0xc00008e220
```

## 3.2.3 一些拓展的结构体知识

声明时赋值

```Go
body2 := Body{
	"tom", 13,
}
```

结构体数组

```Go
bodys := []Body{
	Body{"jack", 12}, Body{"lynn", 18},
}
```

匿名结构体，一般用来存测试用例

```Go
class1 := struct {
	bodys []Body
}{
	[]Body{Body{"jerry", 24}},
}
```

## 3.2.4 小结

通过这篇文章，你应该对 `go` 语言中的 `对象` 有一个直观的体验。

1. 学会如何给结构体绑定方法
2. 了解绑定方法时是否加 `*` 号（指针）的区别
3. 学会声明时赋值、结构体数组、匿名结构体的知识

我们在 `java` 里学习过 `interface` （接口），通过接口定义一系列的函数（标准），实现接口的对象需要实现所有的方法，那 `go` 语言中是否有这种语法呢？我们下次再见！
