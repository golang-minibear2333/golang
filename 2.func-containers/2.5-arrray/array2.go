/*
* @Title:   二维数组
* @Author:  minibear2333
* @Date:    2020-04-04 20:16
* @url:     https://github.com/golang-minibear2333/golang
*/
package main

import (
	"fmt"
	"github.com/golang-minibear2333/golang/tools"
)

func multiArray() {
	//声明二维数组，只要 任意加中括号，可以声明更多维，相应占用空间指数上指
	var arr [3][3]int
	//赋值
	arr = [3][3]int{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
	}
	//访问
	fmt.Printf("arr的第1行，第0列是：%v  \n", arr[1][0])

	fmt.Println("遍历输出")

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[j]); j++ {
			fmt.Printf("%v ", arr[i][j])
		}
		fmt.Println()
	}

	// 不定长数组，但是在go里面被称为切片，不属于数组，这种先提一下，后面的章节会展开
	arr2 := [][]int{{1, 2}, {1, 2, 3}, {1, 2, 3, 4}}
	fmt.Println("输出多维切片")
	tools.PrintMulti2Slice(arr2)

}

