/*
* @Title:   向函数传递数组，引用传递还是值传递？
* @Author:  minibear2333
* @Date:    2020-04-04 20:58
* @url:     https://github.com/golang-minibear2333/golang
*/
package main

func sumArray(arr [5]int) (res int) {
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	arr[0] = 2 // 测试引用传递
	return res
}
