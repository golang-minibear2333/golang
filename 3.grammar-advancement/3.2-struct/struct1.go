/*
* @Title:   结构体
* @Author:  minibear2333
* @Date:    2020-04-05 22:01
* @url:     https://github.com/golang-minibear2333/golang
 */
package main

import "fmt"

type Body struct {
	name string
	age  int
}

func structDemo1() {
	var body Body
	body.name = "coding3min"
	body.age = 12
	fmt.Println(body)

	//声明时赋值
	body2 := Body{
		"tom", 13,
	}
	fmt.Println(body2)

	//结构体切片
	bodys := []Body{
		{"jack", 12}, {"lynn", 18},
	}
	fmt.Println(bodys)

	//匿名结构体，一般用来存测试用例
	class1 := struct {
		bodys []Body
	}{
		[]Body{{"jerry", 24}},
	}
	fmt.Println(class1)
}

func structDemo2() {
	// 结构体指针，事实上和普通变量指针是一样的
	var bodyPtr *Body
	body := Body{"tom", 100}
	bodyPtr = &body
	fmt.Println(*bodyPtr)
	//当然，他可以传递给函数
	funcGetStructPtr(&body)
	fmt.Printf("结构体内容变了：%v \n", *bodyPtr)
	funGetStruct(body)
	fmt.Printf("不使用指针变量，结构体内容不变：%v \n", *bodyPtr)
}
func funcGetStructPtr(bodyPtr *Body) {
	bodyPtr.name = "lisa"
}
func funGetStruct(body Body) {
	body.name = "jj"
}

func main() {
	structDemo1()
	structDemo2()
}
