/*
* @Title:   向函数传递数组，引用传递还是值传递？
* @Author:  pzqu
* @Date:    2020-04-04 20:58
* @url:     https://github.com/pzqu/how_to_code
*/
package main

import "fmt"

func sumArray(arr [5]int) (res int) {
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	arr[0] = 2 // 测试引用传递
	return res
}

func main() {
	arr := [5]int{1, 1, 1, 1, 1}
	fmt.Printf("[5]int{1,1,1,1,1} 的和为：%v \n", sumArray(arr))
	if arr[0]  == 1{
		fmt.Println("函数传数组参数是值传递，会建立副本")
	}else{
		fmt.Println("函数传数组参数是引用传递，直接对原数组操作")
	}
}
