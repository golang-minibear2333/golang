/*
* @Title:   range（范围）
* @Author:  pzqu
* @Date:    2020-04-07 20:31
* @url:     https://github.com/pzqu/how_to_code
*/
package main

import "fmt"

/*
	range 关键字 用于 for 循环中迭代 array、slice、map
	还有 channel 还有字符串、channel
 */
func rangeDemo1() {
	//切片（数组省略）
	nums := []int{1, 2, 3}
	for k, v := range nums {
		fmt.Printf("key: %v , value: %v  \n", k, v)
	}

	//map
	kvs := map[string]string{
		"a":"a",
		"b":"b",
	}
	for k, v := range kvs {
		fmt.Printf("key: %v , value: %v  \n", k, v)
	}

	//字符串
	for k,v := range "hello"{
		fmt.Printf("key: %v , value: %c  \n", k, v) //注意这里单个字符输出的是ASCII码，用 %c 代表输出字符
	}


	// channel （如果不会可以先mark下，详细参考：go的并发特性章节）
	ch := make(chan int, 10)
	ch <- 11
	ch <- 12

	close(ch) // 不用记得关掉,不关掉又没有另一个goroutine存在会死锁哦，可以注释掉这一句体验死锁

	for x := range ch {
		fmt.Println(x)
	}


}

func main(){
	rangeDemo1()
}