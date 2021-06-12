/*
* @Title:   switch和type switch
* @Author:  minibear2333
* @Date:    2020-03-26 09:02
* @url:     https://github.com/golang-minibear2333/golang
*/
package main

import (
	"fmt"
	"math"
)

func ifelse(){
	if 20>0{
		fmt.Println("yes")
	}

	if 20<0{

	}else{
		fmt.Println("no")
	}
}
/*
变量var1可以是任何类型
val1和val2必须是同类型，可以是任意值，可以是常量、变量、表达式，但结果必须是同类型
switch var1{
	case val1:
		...
	case val2:
		...
	default:
		...
}
 */
func switchDemo(name string, number int) {
	switch name {
	case "coding3min":
		fmt.Println("welcome" + name)
	default:
		fmt.Println("403 forbidden:" + name)
		return
	}

	switch {
	case number >= 90:
		fmt.Println("优秀")
	case number >= 80:
		fmt.Println("良好")
	case number >= 60:
		fmt.Println("凑合")
	default:
		fmt.Println("太搓了")
	}
}

//type-switch 用来判断某个interface变量中实际存储的变量类型
//被用于不同版本接口返回json中，属性名一样但是类型有差异,实战：https://github.com/golang-minibear2333/golang/blob/master/golang/medium/json_interface/fixed_json.go
func typeSwitchDemo(x interface{}) int {
	switch t := x.(type) {
	case int:
		return t
	case float64:
		return int(math.Ceil(t))
	}
	return 0
}

func main() {
	ifelse()
	switchDemo("coding3min", 95)
	switchDemo("coding3min", 20)
	switchDemo("tom", 60)

	var x interface{}
	x = 1
	fmt.Println(x)
	fmt.Println(typeSwitchDemo(1.1))
	fmt.Println(typeSwitchDemo(2))
}
