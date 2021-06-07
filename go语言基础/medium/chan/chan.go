/*
* @Title:   通道（chan）的发送与接收
* @Author:  minibear2333
* @Date:    2020-03-27 20:44
* @url:     https://github.com/minibear2333/how_to_code
*/
package main

import (
	"fmt"
	"time"
)

//本节需要先了解goroutine的知识, 示例代码：https://github.com/minibear2333/how_to_code/blob/master/golang/medium/chan/goroutine.go
var c chan int

/*
	通道类型是引用类型，所以他的零值就是nil
	换句话说，只声明但是没有用make函数初始化，改变量的值就是nil
	对于值为nil的通道，不论具体是什么类型，它们所属的接收和发送操作都会永久处于阻塞状态（卡在那里，下面的代码都不会执行）
*/
func init() {
	c = make(chan int, 1)
}

//发送10个数
func send() {
	fmt.Println("send start")
	for i := 1; i <= 10; i++ {
		fmt.Printf("send %v wait \n", i)
		c <- i
		fmt.Printf("send %v end \n", i)
	}
	fmt.Println("send end")
}

//只接收一次
func receive() {
	fmt.Println("receiveChan start")
	var res int
	res = <-c
	fmt.Printf("receive %v \n", res)
	fmt.Println("receiveChan end")
}

//main 本身是一个goroutine
func main() {
	//创建一个新的goroutine，这时它与main goroutine是并发的
	go receive()
	go send()
	//主线程结束的时候强制结束所有的goroutine，所以这里要等待1秒
	fmt.Println("main goroutine sleep start")
	time.Sleep(time.Second)
	/*过程
		创建两个goroutine是，主线程休眠开始
		receive()开始接收，发现通道为空，阻塞receive
		send()发送数字1，尝试发送2，发现通道满了，阻塞send
		receive()发现通道不为空，开始接收，接收到1，只接收一次则关闭go send()
		send()发现通道空了，继续发送2，发送成功，尝试发送3，通道满了，阻塞send
		主线程休眠结束，强制结束了send
	 */
}

