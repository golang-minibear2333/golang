/*
* @Title:   Map(集合)
* @Author:  minibear2333
* @Date:    2020-04-08 22:11
* @url:     https://github.com/golang-minibear2333/golang
 */
package main

import "fmt"

/*
	Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
	Map 是使用 hash 表来实现的,所以 Map 是无序的
*/

func mapDemo1() {
	var m map[string]string
	// 注意了，map 是一种引用类型，初值是nil，必须要申请空间，所有的引用类型都要这么做
	m = make(map[string]string) //如果不相信可以注释掉试试，会报错 panic: assignment to entry in nil map

	// 赋值
	m["name"] = "coding3min"
	m["sex"] = "man"

	// 循环遍历
	for key := range m {
		fmt.Println("key:", key, ",value:", m[key]) // 原来不用Printf也可以完成拼接输出啊！
	}

	// 快速判断元素是否存在
	inMap("name", m)

	// 删除集合元素
	delete(m, "name")
	inMap("name", m)

}

func inMap(key string, m map[string]string) {
	// 快速判断元素是否存在
	if value, ok := m[key]; ok {
		fmt.Println(key, "存在，值为：", value)
	} else {
		fmt.Println(key, " 不存在")
	}
}

// map 的拷贝
func copyMap() {
	fmt.Println("map 的拷贝")
	originalMap := make(map[string]int)
	originalMap["one"] = 1
	originalMap["two"] = 2

	// Create the target map
	targetMap := make(map[string]int)

	// Copy from the original map to the target map
	for key, value := range originalMap {
		targetMap[key] = value
	}
	fmt.Println(targetMap)
}

// 深度拷贝
func copyDeepMap() {
	fmt.Println("深度拷贝")
	originalMap := make(map[string]*int)
	var num int = 1
	originalMap["one"] = &num

	// Create the target map
	targetMap := make(map[string]*int)

	// Copy from the original map to the target map
	for key, value := range originalMap {
		var tmpNum int = *value
		targetMap[key] = &tmpNum
	}
}

type Person struct {
	age int
}

// 注意修改map中value内部数据
func updateMapValue() {
	originalMap := make(map[string]Person)
	originalMap["minibear2333"] = Person{age: 26}
	// 放开以下注释体验报错
	//originalMap["minibear2333"].age = 5

	originalPointMap := make(map[string]*Person)
	originalPointMap["minibear2333"] = &Person{age: 26}
	originalPointMap["minibear2333"].age = 5
}
func main() {
	mapDemo1()
	// 小熊的话：能用数组别用map，数组快占用空间小
	// 但是要在保证快速开发的情况下再考虑用数组优化
	copyMap()
	copyDeepMap()

}
