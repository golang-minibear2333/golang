/*
* @Title:  声明【常量】的各种方式
* @Author: pzqu
* @Date:   2020-03-24 22:09
*/
package main

import "fmt"

func constVariable() {
	// 定义方式 const identifier [type] = value
	// 约定常量全大写表示
	const A int = 1
	const B = 1

	const C, D, E = 1, 1, 1

	//枚举,一般枚举绑定到struct类型
	const (
		Success = 0
		UnKonw  = 1
		Error   = 2
	)

	//特殊常量，被认为是可以被编译器修改的常量
	//const 出现时被重置为0，每出现一次自动加1
	const (
		F = iota
		G = iota
		H = iota
	)
	fmt.Println(F, G, H)
	// 可以简写
	const (
		I = iota
		J
		K
	)
	fmt.Println(I, J, K)
}

func main() {
	constVariable()
}
