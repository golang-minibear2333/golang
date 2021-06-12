/*
* @Title:   指针
* @Author:  minibear2333
* @Date:    2020-04-05 21:01
* @url:     https://github.com/golang-minibear2333/golang
*/
package main

import "fmt"

// 其实我们在go语言中的类，这一节已经稍微带了一下指针了
// how_to_code/base/easy/function/go_class.go

func main() {
	// 和c++中的指针用法一样(看不懂可以给我提issue: https://github.com/golang-minibear2333/golang/issues)
	var a int
	fmt.Printf("a 的地址是：%p \n", &a)
	//声明 变量名 + 指针类型 , 命令规则以ptr结尾
	var ptr *int /* 指向整型*/
	// var fp *float32 /* 指向浮点型 */
	ptr = &a // 变量内部存的值是普通类型，指针内部存的值是地址
	fmt.Printf("ptr 存的值是：%p \n", ptr)

	if a == *ptr {
		fmt.Println("a == *ptr ，存的就是a的地址，ptr的指向*ptr肯定就是a本身了")
	}

	// 最后要注意，声明未初始化，初始化值为 nil ，在某些地方会出大问题

}
