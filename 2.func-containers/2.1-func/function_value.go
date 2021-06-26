// package main 函数当作变量使用，当做 参数传递
package main

import "fmt"

//函数作为值使用
func functionValue(a, b int, do func(int, int) int) {
	fmt.Println(do(a, b))
}

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func main() {
	functionValue(1, 1, add)
	functionValue(1, 1, sub)
	//匿名函数
	functionValue(1, 1, func(i1 int, i2 int) int {
		return i1 * i2
	})
	// 实际的使用
	// 你可以参考函数测速例子：https://github.com/golang-minibear2333/golang/blob/master/golang/easy/string/append_string.go
	// 还有你可以传filter函数做过滤，mapping做映射等实际的用法
	// 有时候也可以作为排序递增，递减的依据
}
