/*
* @Title:  panic详解
* @Author: minibear2333
* @Date:   2020-05-28 13:24
*/
package main

import (
	"fmt"
)

func divisionIntRecover(a, b int) (ret int) {
	defer func() {
		if err := recover(); err != nil {
			// 打印错误，关闭资源，退出此函数
			fmt.Println(err)
			ret = -1
		}
	}()

	return a / b
}

func genPanic() {
	defer fmt.Println("关闭文件句柄")

	panic("人工创建的运行时异常")
}

func main() {

	var res int
	datas := []struct {
		a int
		b int
	}{
		{2, 0},
		{2, 2},
	}

	for _, v := range datas {
		if res = divisionIntRecover(v.a, v.b); res == -1 {
			continue
		}
		fmt.Println(v.a, "/", v.b, "计算结果为：", res)
	}

}
