/*
* @Title:  线程安全的map
* @Author: minibear2333
* @Date:   2020-05-06 22:06
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.Map
	scene.Store("name", "coding3min")
	scene.Store("age", 11)

	v, ok := scene.Load("name")
	if ok {
		fmt.Println(v)
	}
	v, ok = scene.Load("age")
	if ok {
		fmt.Println(v)
	}

	scene.Delete("age")

	scene.Range(func(key, value interface{}) bool {
		fmt.Println("key:",key,",value:",value)
		return true
	})
}
