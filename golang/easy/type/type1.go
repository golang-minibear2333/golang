/*
* @Title:   基本类型转换
* @Author:  minibear2333
* @Date:    2020-04-09 12:41
* @url:     https://github.com/minibear2333/how_to_code
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
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
