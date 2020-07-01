/*
* @Title:   测速
* @Author:  minibear2333
* @Date:    2020-03-17 00:29
* @url:     https://github.com/minibear2333/how_to_code
*/
package utils

import (
	"fmt"
	"time"
)

func doWhat(){

}
func SpeedTime(handler func() , funcName string) {
	t := time.Now()
	handler()
	elapsed := time.Since(t)
	fmt.Println(funcName+"spend time:", elapsed)
}
func main() {
	SpeedTime(doWhat, "doWhat")
}

