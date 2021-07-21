/*
* @Title:   select语法
* @Author:  minibear2333
* @Date:    2020-03-27 20:23
* @url:     https://github.com/golang-minibear2333/golang
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	selectDemo()
}

func init() {
	c2 := make(chan int, 1)
	go func() {
		for i := 1; i <= 10; i++ {
			c2 <- 1
		}
		close(c2)
	}()
	/*
		select 语句会随机执行一个case，如果没有case可以运行，就会一直阻塞，直到有case可以运行
		case 必须是一个通信操作，要么是发送，要么是接收
		如果有default体，就只运行default，其他全部忽略
	*/
}

func selectDemo() {
	countAdd, countSub := 0, 0
	// 形式1：异步式，这里会接收一个值，然后sendDemo()会卡在插入3之前
	select {
	case _, ok := <-c2:
		if ok {
			countAdd++
			fmt.Printf("c2 <- 1 , countAdd: %v\n", countAdd)
		} else {
			fmt.Println("close")
			break
		}

	case _, ok := <-c2:
		if ok {
			countSub++
			fmt.Printf("c2 <- 1 , sub count: %v\n", countSub)
		} else {
			fmt.Println("close")
			break
		}

	}

	//阻塞式,一个返回值，如果c2里面是空的就一直阻塞了，这里会接收一个值，然后sendDemo()会卡在插入4之前
	select {
	case _ = <-c2:
		countAdd++
		fmt.Printf("c2 <- 1 , countAdd: %v\n", countAdd)
	case _ = <-c2:
		countSub++
		fmt.Printf("c2 <- 1 , sub count: %v\n", countSub)
	}
	time.Sleep(time.Second)

	//阻塞时，运行default
	select {
	case c2 <- -1:
		countAdd++
		fmt.Printf("c2 <- 1 , countAdd: %v\n", countAdd)
	default:
		fmt.Println("c2 chan is full!! can't insert number")
	}
	time.Sleep(time.Second)

	//不阻塞时，不运行default，运行任意一个可以运行的case
	select {
	case _ = <-c2:
		countAdd++
		fmt.Printf("c2 <- 1 , countAdd: %v\n", countAdd)
	default:
		fmt.Println("if case ok, default can't  run")
	}
	//PS: 以上就是用法，但用法是这个用法，但是实在想不出具体的使用场景
}
