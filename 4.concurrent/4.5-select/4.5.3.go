package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func A() int {
	fmt.Println("start A")
	time.Sleep(1 * time.Second)
	fmt.Println("end A")
	return 1
}
func B() int {
	fmt.Println("start B")
	time.Sleep(1 * time.Second)
	fmt.Println("end B")
	return 2
}

func SpeedTime(handler func()) {
	t := time.Now()
	handler()
	elapsed := time.Since(t)
	// 利用反射获得函数名
	funcName := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
	fmt.Println(funcName+"spend time:", elapsed)
}
func lee() {
	ch, done := make(chan int), make(chan struct{})
	defer close(ch)
	go func() {
		select {
		case ch <- A():
		case ch <- B():
		case <-done:
		}
	}()
	done <- struct{}{}
}
func main() {
	SpeedTime(lee)
}
