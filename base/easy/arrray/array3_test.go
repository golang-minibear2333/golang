package main

import (
	"fmt"
	"testing"
)

func TestSumArray(t *testing.T){
	arr := [5]int{1, 1, 1, 1, 1}
	fmt.Printf("[5]int{1,1,1,1,1} 的和为：%v \n", sumArray(arr))
	if arr[0]  == 1{
		fmt.Println("函数传数组参数是值传递，会建立副本")
	}else{
		fmt.Println("函数传数组参数是引用传递，直接对原数组操作")
	}
}