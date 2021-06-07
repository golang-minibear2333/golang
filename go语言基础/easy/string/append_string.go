/*
* @Title:   快速拼接字符串
* @Author:  minibear2333
* @Date:    2020-03-17 00:12
* @url:     https://github.com/minibear2333/how_to_code
*/
package main

import (
	"bytes"
	"fmt"
	"time"
)

var S string

func init() {
	S = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" +
		"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" +
		"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" +
		"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" +
		"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}

/*
	go语言基础 里面的字符串都是不可变的，每次运算都会产生一个新的字符串，
	所以会产生很多临时的无用的字符串，不仅没有用，还会给 gc 带来额外的负担，
	所以性能比较差
 */
func appendStr() (resStr string) {
	resStr = ""
	for i := 0; i < len(S); i++ {
		resStr += string(S[i])
	}
	return resStr
}

func appendStrQuick() (string) {
	var res bytes.Buffer
	for i := 0; i < len(S); i++ {
		res.WriteString(string(S[i]))
	}
	return res.String()
}

func speedTime(handler func() (string), funcName string) {
	t := time.Now()
	handler()
	elapsed := time.Since(t)
	fmt.Println(funcName+"spend time:", elapsed)
}
func main() {
	speedTime(appendStr, "appendStr")
	speedTime(appendStrQuick, "appendStrQuick")
}
