// package main 匿名函数
package main

import "fmt"

func functionValue(a, b int, do func(int, int) int) {
	fmt.Println(do(a, b))
}

func main(){
	functionValue(1,2,func(a,b int) int{
		return a+b })
	//使用匿名函数的方法调用他 实现匿名减函数
	functionValue(1,2,func(a,b int) int{
		return a-b })

	f := func(i int) {
		fmt.Println(i)
	}

	f(1)
}
