/*
* @Title:   递归函数
* @Author:  pzqu
* @Date:    2020-04-09 12:35
* @url:     https://github.com/pzqu/how_to_code
*/
package main

import "fmt"

// 就是自己调用自己(注意，在复杂度达到一定程序时不建议使用，由于递归要维护很多函数栈，会占用时间和空间)
// 递归实现阶乘
func fectorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * fectorial(x-1)
}

func main() {
	fmt.Printf("8 的阶乘为 %v", fectorial(8))
}
