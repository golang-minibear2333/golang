package main

import (
	"fmt"
	"strconv"
)

// 类型转换
func convertType() {
	var aInt int = 17

	// 一般用这种方式强制转
	fmt.Printf("转float64 %f  \n", float64(aInt))
	fmt.Printf("转string %v  \n", strconv.Itoa(aInt))
	fmt.Printf("转float64 %f  \n", float64(aInt))

	// 各种类型转字符串
	resString := fmt.Sprintf("%d %v %v", 1, "coding3min", true)
	fmt.Println(resString)

	// string  to bytes
	resBytes := []byte("asdfasdf")
	// bytes to string
	resString = string(resBytes)
	fmt.Println(resString)

}
