/*
* @Title:   interface 类型（接口）
* @Author:  minibear2333
* @Date:    2020-04-09 12:48
* @url:     https://github.com/golang-minibear2333/golang
*/
package main

import "fmt"

/*
	同 java 一样，可以把一堆有共性的方法定义在里面
	但是比 java 灵活的是，不需要显式实现接口，你可以自己控制实现哪些方法。
	但必须实现了所有接口的方法才算是实现了这个接口，并用下面这样的格式来实现他
		接口实例 = new(类型)
*/
type humanInterface interface {
	eat() string
	play() string
}

type man struct {
	name string
}

func (p man) eat() string {
	return "eat banana"
}

func (p man) play() string {
	return "play game"
}

func main() {
	var human humanInterface
	human = new(man)
	fmt.Println(human.eat())
	fmt.Println(human.play())
}
