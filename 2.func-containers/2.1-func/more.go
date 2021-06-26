// package main 值传递和引用传递
package main

import "fmt"

//函数参数

//值传递
func noChange(a, b int) {
	tmp := a
	a = b
	b = tmp
}

// 引用传递，参数加*号代表指针
func change(a,b *int){
	tmp := *a
	*a = *b
	*b = tmp
}

func main() {
	a, b := 1, 2
	fmt.Printf("原值 a:%v,b:%v \n", a, b)
	noChange(a, b)
	//值传递，并没有修改原值
	fmt.Printf("值传递后 a:%v,b:%v \n", a, b)
	//引用传递，&就是c中的取地址
	change(&a,&b)
	fmt.Printf("引用传递后 a:%v,b:%v \n", a, b)

}
