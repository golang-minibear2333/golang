/*
* @Title:   协程（goroutine）
* @Author:  minibear2333
* @Date:    2020-03-27 21:05
* @url:     https://github.com/golang-minibear2333/golang
*/
package main

import (
	"fmt"
	"time"
)

func quickFun(){
	fmt.Println("maybe you can's see me!")
}


func main(){
	go quickFun() //创建了一个goroutine（语言级别的协程）
	//然后协程和main主线程同时运行
	fmt.Println("hey")

	go func() {
		fmt.Println("hello ")
	}()

	time.Sleep(time.Second) //main运行结束会暴力终止所有协程，所以这里先等待1秒


}
