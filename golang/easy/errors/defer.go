/*
* @Title:  defer语句
* @Author: pzqu
* @Date:   2020-05-28 13:14
*/
package main

import "fmt"

func main() {

	defer fmt.Println("see you next time!")

	defer fmt.Println("close all connect")

	fmt.Println("hei boy")
}
