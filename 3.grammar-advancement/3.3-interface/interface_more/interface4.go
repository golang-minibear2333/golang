package main

import "fmt"

// 声明接口类型
type humanInterface interface {
	//name string 很可惜，接口内是不可以声明成员变量的,用接口的方式主要目标是函数的多态，提倡在类函数内使用成员变量
	eat() string
	play() string
}
// 类1
type man struct {
	name string
}

func (p man) eat() string {
	return p.name + "eat banana"
}

func (p man) play() string {
	return p.name + "play game"
}
// 类2
type woman struct {
	name string
}

func (p woman) eat() string {
	return p.name + "eat rice"
}

func (p woman) play() string {
	return p.name + "watch TV"
}

// 多态
func humanDoWhat(p humanInterface) {
	fmt.Println(p.eat())
	fmt.Println(p.play())
}

func main() {
	w := woman{"lisa"}
	m := man{"coding3min"}
	// 多态的含义就是不需要修改函数，只需要修改外部实现
	// 同一个接口有不同的表现
	humanDoWhat(w)
	humanDoWhat(m)
}
