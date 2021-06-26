package main

import "fmt"

// 循环语句的多种形式、死循环、break/continue
func demo() {
	a, b := 1, 5
	nums := []int{1, 2, 3, 4, 5, 6}
	//方式1
	for i := 0; i < len(nums); i++ {
		fmt.Println(i)
	}
	//方式2
	for a < b {
		fmt.Println(a)
		a++
	}
	//方式3
	for index, value := range nums {
		fmt.Printf("key: %v , value: %v \n", index, value)
	}
	//当然，你可以把方式3中index去掉
	for _, v := range nums {
		fmt.Printf("value: %v \n", v)
	}
	//也可以只返回一个
	for i := range nums {
		fmt.Printf("value: %v \n", nums[i])
	}

	i := 0
	//死循环,就是什么循环条件也不要加
	for {
		fmt.Printf("死循环测试 %v \n", i)
		i++
		if i > 5 {
			fmt.Println("满足终止条件，退出")
			break //直接跳出循环
		}
		if i == 3 {
			continue //会直接跳过执行后面的语句
		}
		fmt.Printf("死循环测试,第%v次跑到循环结尾\n", i)
	}
	//goto 语法也支持，但是会打破结构，依然不建议用，感兴趣自己下来了解

}

func main() {
	demo()
}
