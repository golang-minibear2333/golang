// package main 函数的简单使用
package main

import "fmt"

//一个函数
func functionSample() {

}

//一个参数的函数
func functionParam(num int) {
}

//多个参数的函数
func functionParams(a, b int, c string) {
}

//一个返回值
func funcReturnOne() int {
	return 1
}

//多个返回值
func funReturnMany() (int, int) {
	return 1, 2
}

//返回值有名称
func funReturnName() (res int) {
	res = 1 + 1
	return
}

func main() {
	//接收多个返回值
	a, b := funReturnMany()
	fmt.Println(a + b)
}
