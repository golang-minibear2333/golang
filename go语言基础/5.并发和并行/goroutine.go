package main

import "fmt"

func Add(a, b int) {
	c := a + b
	fmt.Println(c)
}

func main() {
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}
}
