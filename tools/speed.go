package tools

import (
	"fmt"
	"time"
)

func doWhat() {

}
func SpeedTime(handler func(), funcName string) {
	t := time.Now()
	handler()
	elapsed := time.Since(t)
	fmt.Println(funcName+"spend time:", elapsed)
}
func main() {
	SpeedTime(doWhat, "doWhat")
}
