// 数组定义赋值与遍历
package main

import (
	"fmt"

	"github.com/golang-minibear2333/golang/tools"
)

func arrayDefine() {
	//声明定长数组
	var a1 [10]int
	//声明不定长数组(切片)
	var a2 []int

	//初始化数组
	var b1 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	var b2 = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	/* 为数组 a1 初始化元素 */
	for i := 0; i < len(a1); i++ {
		a1[i] = i + 100 /* 设置元素为 i + 100 */
	}
	/* 输出每个数组元素的值 */
	for j := 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, a1[j])
	}

	//忽略未使用错误
	tools.IgnoreUnused(a1)
	tools.IgnoreUnused(a2)
	tools.IgnoreUnused(b1)
	tools.IgnoreUnused(b2)
}

func main() {

}
