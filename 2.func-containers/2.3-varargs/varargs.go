// package main 可变参数
package main

import (
	"fmt"
	"strconv"
)

func sum(t ...int) (res int) {
	for _, v := range t {
		res += v
	}
	return res
}


func sumNumSample(t ...interface{}) (res float64){
	for _,tmp := range t{
		switch v :=tmp.(type) {
		case int:
			res += float64(v)
		case float64:
			res+= v
		case float32:
			res += float64(v)
		}
	}
	return res
}

func sumNum(t ...interface{}) (res float64) {
	for _, tmp := range t {
		switch v := tmp.(type) {
		case int, int8, int16,
		int32, int64, uint,
		uint8, uint16, uint32,
		uint64, float32, float64,
		complex64, complex128:
			convertStr := fmt.Sprintf("%v", v)
			convertFloat64, _ := strconv.ParseFloat(convertStr, 64)
			res += convertFloat64
		}
	}
	return res
}

func main() {
	fmt.Println("原生的可变参数例子")
	fmt.Println("a", "b", "c", "d")

	fmt.Println("简单的可变参数")
	fmt.Println(sum(1, 2, 3, 4, 5))

	fmt.Println("不定参数，不定类型")
	fmt.Println(sumNumSample(1, 2.1, "asd", true))
	fmt.Println(sumNum(1, 2.1, "asd", true))
	// 不定参数格式化输出
	// func Printf(format string, a ...interface{}) (n int, err error) {

	// 其他可用地方，字符串的拼接，不定参数运算

}
