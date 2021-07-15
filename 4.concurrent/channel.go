package main

import "fmt"

func main() {
	a := make(chan int)
	a <- 1   //将数据写入channel
	z := <-a //从channel中读取数据
	fmt.Println(z)
}
