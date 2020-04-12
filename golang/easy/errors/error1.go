/*
* @Title:   错误处理初识
* @Author:  pzqu
* @Date:    2020-04-11 21:46
* @url:     https://github.com/pzqu/how_to_code
*/
package main

import (
	"errors"
	"fmt"
)

func errorDemo() (int,error){
	err := errors.New("这是一个自定义错误")

	printErrMessage(err)

	// 组合错误信息

	errorStr := "产生了一个%v错误"
	errMessage := fmt.Sprintf(errorStr, "吃太多")
	err = errors.New(errMessage)

	printErrMessage(err)

	// 组合错误信息改良（一行）
	err = fmt.Errorf(errorStr, "喝太多")

	printErrMessage(err)

	return 0,err

}

func printErrMessage(err error) {
	// 判断上一个语句返回，是否包含错误的过程有个学名叫：卫述语句
	// go 语言会有一大堆这种语句，在go2中会被 "check..handle" 语句修复
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//根据Golang的约定，每个可能导致错误的函数都将error其作为最后一个返回值，码农有责任在每一步都正确处理它
	_,err := errorDemo()
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("没有错误返回，继续执行")
}
