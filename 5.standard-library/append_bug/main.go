package main

import "fmt"

func main() {
	var a = make([]int, 0)
	for i := 0; i < 3; i++ {
		a = append(a, i)
	}
	b := append(a, 1)
	c := append(a, 2)
	d := append(a, 3)
	fmt.Println(b[3], c[3], d[3])
}
