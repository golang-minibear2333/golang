/*
* @Title:  
* @Author: minibear2333
* @Date:   2020-05-06 22:32
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	slice2 := []int{0, 0, 0, 1, 2, 3}
	slice3 := slice2[0:1]
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice3), cap(slice3), slice3)
	slice3[0] = 1
	slice2[0] = 2
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice2), cap(slice2), slice2)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice3), cap(slice3), slice3)

	fmt.Println("修改后")
	slice2 = []int{0, 0, 0, 1, 2, 3}
	slice3 = make([]int, 1, 1)
	copy(slice3, slice2[0:1])
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice3), cap(slice3), slice3)
	slice3[0] = 1
	slice2[0] = 2
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice2), cap(slice2), slice2)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice3), cap(slice3), slice3)

	// 排序函数
	slice2 = []int{0, 3, 0, 1, 2, 0}
	sort.Ints(slice2)
	fmt.Println(slice2)

}
