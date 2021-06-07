package main

import "fmt"

// 先创建一个结构体 Animal，再以 Animal 为基类创建结构体 Cat 和 Dog,
// 这两个子类分别实现接口 behavior 的 eat() 和 run() 方法。

type Animal struct {
}
type Cat struct {
	Animal
}
type Dog struct {
	Animal
}
type behavior interface {
	eat()
	run()
}

func (c Cat) eat() {
	fmt.Println("cat eat")
}

func (c Cat) run() {
	fmt.Println("cat run")
}

func (d Dog) eat() {
	fmt.Println("dog run")
}
func (d Dog) run() {
	fmt.Println("dog run")
}
func main(){
	list := make([]behavior,2)
	list[0] = Cat{}
	list[1] = Dog{}
	for _,v := range list{
		v.run()
		v.eat()
	}
}
