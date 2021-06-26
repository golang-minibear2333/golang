package main

import "fmt"

// 各种常量定义方式
func constVariable() {
	// 定义方式 const identifier [type] = value
	// 约定常量全大写表示
	const A int = 1
	const B = 1

	const C, D, E = 1, 1, 1

	//枚举
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
