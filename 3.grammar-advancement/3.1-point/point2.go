/*
* @Title:   指针扩展知识
* @Author:  minibear2333
* @Date:    2020-04-05 21:19
* @url:     https://github.com/golang-minibear2333/golang
*/
package main

import "fmt"

//多维度指针
func point21() {
	var a int
	var ptr *int
	var pptr **int
	var ppptr ***int

	ptr = &a
	pptr = &ptr
	ppptr = &pptr
	fmt.Printf("a的地址：%p \n", &a)
	fmt.Printf("ptr存的地址：%p \n", ptr)
	fmt.Printf("pptr存的地址的指向：%p \n", *pptr)
	fmt.Printf("ppptr存的地址的指向的指向：%p \n", **ppptr)

}

//把地址当作参数
func point22() {
	a := 100
	b := 200

	//操作地址，不需要返回
	swap(&a, &b)
	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}

func main() {
	point21()
	point22()
}
