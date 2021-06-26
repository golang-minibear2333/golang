/*
* @Title:   这不是接口
* @Author:  minibear2333
* @Date:    2020-04-09 12:59
* @url:     https://github.com/golang-minibear2333/golang
*/
package main

type dogInterface interface {
	eat() string
	play() string
}

type dog1 struct {
	name string
}

func (d dog1) eat() string {
	return "Eat dog food"
}

func main() {
	// 报错：Cannot use 'new(dog1)' (type *dog1) as type dogInterface in assignment
	// Type does not implement 'dogInterface' as some methods are missing: play() string more...
	//var dog dogInterface
	// dog = new(dog1) 放开注释查看报错
}
