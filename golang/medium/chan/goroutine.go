/*
* @Author: pzqu
* @Date:   2020-03-27 21:05
*/
package main

import "fmt"

func quickFun(){
	fmt.Println("maybe you can's see me!")
}

//创建了一个goroutine（语言级别的协程）
//然后协程和main主线程同时运行
//main运行结束会暴力终止所有协程，有概率quickFun中的输出还没来得及执行
//可以多跑一次看看效果
func main(){
	go quickFun()
	fmt.Println("hey")
}
