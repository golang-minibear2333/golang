package tools

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func doWhat() {

}
func SpeedTime(handler func()) {
	t := time.Now()
	handler()
	elapsed := time.Since(t)
	// 利用反射获得函数名
	funcName := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
	fmt.Println(funcName+"spend time:", elapsed)
}
func main() {
	SpeedTime(doWhat)
}
