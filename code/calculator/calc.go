package main

import "fmt"

func add(x,y int) int{
	return x+y
}
func subtrace(x,y int) int{
	return x-y
}
func multiply(x,y int) int{
	return x*y
}
func divide(x,y int) int{
	return x/y
}
func main() {
	var x,y = 1,1
	fmt.Printf("%d + %d = %d\n",x,y,add(x,y))
	fmt.Printf("%d - %d = %d\n",x,y,subtrace(x,y))
	fmt.Printf("%d * %d = %d\n",x,y,multiply(x,y))
	fmt.Printf("%d / %d = %d\n",x,y,divide(x,y))
}
