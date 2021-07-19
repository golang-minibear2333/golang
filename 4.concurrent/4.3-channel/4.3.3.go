package main

import "fmt"

// 阻塞
func pendingForever() {
	a := make(chan int)
	a <- 1   //将数据写入channel
	z := <-a //从channel中读取数据
	fmt.Println(z)
}

// 正常使用
func normal() {
	fmt.Println("正常使用")
	chanInt := make(chan int)
	go func() {
		chanInt <- 1
	}()

	res := <-chanInt
	fmt.Println(res)
}

// 标准用法
func standard() {
	fmt.Println("标准用法")
	chanInt := make(chan int)
	go func() {
		defer close(chanInt)
		var produceData = []int{1, 2, 3}
		for _, v := range produceData {
			chanInt <- v
		}
	}()
	for v := range chanInt {
		fmt.Println(v)
	}
}
func main() {
	// 死锁，放开注释体验
	//pendingForever()
	normal()
	standard()
}
