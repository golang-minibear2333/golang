/*
* @Title:  声明【变量】的各种方式
* @Author: pzqu
* @Date:   2020-03-23 20:02
*/
package main

import "fmt"

func sampleDefineVariable() {
	// var identifier type
	var name string
	name = "s"

	//根据赋值自动判断类型
	var p = name

	//直接声明并赋值（必须是初次声明才有冒号）
	p2 := "as"

	//多变量声明,不赋值自动赋值为0，比如d e f
	var a, b, c = 1, 2, 3
	var d, e, f int
	h, i, j := 1, 2, 3

	//类型不同的多个变量，难看的要死
	var (
		k int
		l string
	)

	//这样好看
	var m, n, o = "a", 1, true

	fmt.Println(name + p + p2 + l + m)
	fmt.Println(a + b + c + d + e + f + h + i + j + k + n)
	fmt.Println(o)
}
