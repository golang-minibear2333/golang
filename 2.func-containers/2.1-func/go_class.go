// package main 函数方法(go中定义一个类)
package main

import "fmt"

// 在go语言中定义类的 形式就 是 绑定在struct类型中(关于struct可以参考how_to_code/base/easy/struct/struct1.go)
type people struct {
	name string
}

func (p people) toString() {
	fmt.Println(p.name)
	fmt.Printf("p的地址 %p \n", &p)
}

func (p *people) sayHello() {
	fmt.Printf("Hello! %v \n", p.name)
	fmt.Printf("*p的地址 %p \n", p)
}

func main() {
	p1 := people{"coding3min"}
	p1.toString()
	p1.sayHello()
	p2 := &people{"tom"}
	p2.toString()
	p2.sayHello()
	//所以用不用指针在使用上没有区别
	fmt.Printf("p1的地址 %p \n", &p1)
	fmt.Printf("p2的地址 %p \n", p2)
	//通过输出地址，你会发现，用指针不会传递额外的对象
}
