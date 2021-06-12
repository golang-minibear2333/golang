/*
* @Title:   匿名函数和闭包
* @Author:  minibear2333
* @Date:    2020-04-01 23:21
* @url:     https://github.com/golang-minibear2333/golang
*/
package main

import (
	"fmt"
	"time"
)


//函数作为值使用
func functionValue2(a, b int, do func(int, int) int) {
	fmt.Println(do(a, b))
}

/*
	匿名函数和闭包
	好处：可以减少全局变量防止变量污染
	坏处：延长了局部变量和函数的生命周期，增加了gc的压力
 */
// 匿名函数，其实之前我们在函数变量那一节用过一次
func noNameFunc() {
	f := func(i int) {
		fmt.Println(i)
	}

	f(1)
}

// 闭包简单实现
func closureSample() func() {
	count := 0
	return func() {
		count ++
		fmt.Printf("调用次数 %v \n", count)
	}
}

func main() {
	// 匿名函数的例子
	functionValue2(1,2,func(a,b int) int{
		return a+b })

	noNameFunc()
	// 声明了两个函数变量
	c1, c2 := closureSample(), closureSample()
	c1()
	c1()
	c1()
	// 你会发现c2又从1开始输出，因为两个函数的变量是独立使用的
	c2()
	c2()
	// 因为各个函数是独立使用一套自己的内部变量，互相不影响，所以闭包可以当测试用例使用
	// 传入不同的实现，不用定义全局变量

	//  闭包形式2，立即执行函数
	func() {
		// to do something
	}()

	// 闭包存在一个问题,如果函数内的闭包是延迟调用的，比较go创建了goroutine，或者defer
	// 这个时候因为外部变量（例如i）已经改变为3了，for循环已经执行完了才执行闭包
	// 会导致此时闭包使用外部变量的时候，出现下面的情况(输出3个3)
	for i := 0; i < 3; i++ {
		fmt.Printf("第一次 i 产生变化中 %v \n", i)
		go func() {
			fmt.Printf("第一次输出： %v\n", i)
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("")
	//这种仅仅是在匿名函数延迟使用的时候会发生
	//解决办法，创建副本，可以给匿名函数加一个参数，传值过来自动生成副本
	for i := 0; i < 3; i++ {
		fmt.Printf("第二次 i 产生变化中 %v \n", i)
		go func(tmp int) {
			fmt.Printf("第二次输出： %v\n", tmp)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("")
	//第二种创建副本的形式
	for i := 0; i < 3; i++ {
		fmt.Printf("第三次 i 产生变化中 %v \n", i)
		tmp := i
		go func() {
			fmt.Printf("第三次输出： %v\n", tmp)
		}()
	}
	time.Sleep(time.Second)

}
