/*
* @Author: pzqu
* @Date:   2020-03-17 00:29
*/
package main

import (
	"fmt"
	"time"
)

func doWhat(){

}
func speedTime(handler func() , funcName string) {
	t := time.Now()
	handler()
	elapsed := time.Since(t)
	fmt.Println(funcName+"spend time:", elapsed)
}
func main() {
	speedTime(doWhat, "doWhat")
}

