/*
* @Title:   切片的长度与容量，len cap append copy
* @Author:  minibear2333
* @Date:    2020-04-06 00:45
* @url:     https://github.com/golang-minibear2333/golang
*/
package main

import "fmt"

func sliceDemo2() {
	slice1 := []int{1, 2, 3}
	printSlice(slice1)
	// slice1[len(slice1)] = 4 //  虽然是不定长的切片，但是一旦赋值容量就固定了，这里会报错

	//指定容量
	slice1 = make([]int, 3, 5) // 3 是长度 5 是容量
	printSlice(slice1)

	// slice1[len(slice1)] = 4 // 还是会报错，虽然容量是 5 ，但是数组长度是3，这里是以长度为准，而不是容量

	//使用append来追加元素，append很常用
	slice1 = append(slice1, 4)
	printSlice(slice1)

	slice1 = append(slice1, 5)
	slice1 = append(slice1, 6) // 到这里长度超过了容量，容量自动翻倍为 5*2
	printSlice(slice1)

	// 上面容量自动翻倍的过程可以看作和下面一致
	slice1 = make([]int, 3, 5) // 3 是长度 5 是容量
	slice1 = append(slice1, 4)
	slice1 = append(slice1, 5)

	// 长度不变，容量自动翻倍为 5*2
	slice2 := make([]int, len(slice1), (cap(slice1))*2)

	/* 拷贝 slice1 的内容到 slice2 */
	copy(slice2, slice1) // 注意是后面的拷贝给前面
	slice2 = append(slice2, 6)
	printSlice(slice2)

}

func printSlice(x []int) {
	// 使用len()测长度，使用cap()测容量
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func main() {
	sliceDemo2()
}
