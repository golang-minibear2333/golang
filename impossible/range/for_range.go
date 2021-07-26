package main

import "fmt"

type Test struct {
	Index int
	Num   int
}

func main() {
	var t []Test
	t = append(t, Test{Index: 1, Num: 1})
	t = append(t, Test{Index: 2, Num: 2})

	// 实际上没有成功修改t.Num，因为是副本复制
	for _, v := range t {
		v.Num += 100
	}

	for _, v := range t {
		// 输出
		// 1 1
		// 2 2
		fmt.Println(v.Index, v.Num)
	}
	// map 也不能这么搞，实际上都是复制
	m := make(map[int]Test)
	m[0] = Test{Index: 1, Num: 1}
	m[1] = Test{Index: 2, Num: 2}
	for _, v := range m {
		v.Num += 100
	}
	for _, v := range m {
		// 输出(可以乱序)
		// 1 1
		// 2 2
		fmt.Println(v.Index, v.Num)
	}

	//两个办法，用下标（map也一样）
	for i := range t {
		t[i].Num += 100
		fmt.Println(t[i].Num)
		// 输出(可以乱序)
		// 101 102
	}

	// 用指针
	var t2 []*Test
	t2 = append(t2, &Test{Index: 1, Num: 1})
	t2 = append(t2, &Test{Index: 2, Num: 2})

	for k, v := range t2 {
		v.Num += 100
		fmt.Println(t2[k].Num)
		// 输出(可以乱序)
		// 101 102
	}

	// https://studygolang.com/articles/25094?fr=sidebar
}
