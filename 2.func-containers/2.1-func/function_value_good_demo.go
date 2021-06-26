// package main 函数当作变量使用，当做 参数传递的分页实践
package main

import (
	"errors"
	"fmt"
)

/*
	再举一个很棒的例子
	假如你在某个接口是分页的，
	需要自己判断页码
 */

type class struct {
	names []string
	index int //当前页码
}

//一个很棒的例子
func (c *class) forEach(hander func([]string) (bool, error)) error {
	//死循环，用来翻页
	for {
		//如果第一个元素是空的，尝试翻页
		if c.index < 0 {
			if err := c.nextClass(); err != nil {
				return err
			}
		}

		//遍历每个单元的值，传给函数使用
		ok, err := hander(c.names)
		if err != nil {
			return err
		}
		if !ok {
			return nil
		}

		//翻页
		if err := c.nextClass(); err != nil {
			return err
		}
		if len(c.names) == 0 {
			return nil
		}
	}
}

//翻页
func (c *class) nextClass() (err error) {
	if c.index < 0 {
		c.index = 0
	} else {
		c.index ++
	}
	//判断是否有下一页
	if c.index < len(tables) {
		c.names = tables[c.index].names
	} else {
		//翻到底了
		c.names = []string{}
	}
	return nil
}

//测试数据
var tables []class

func init() {
	tables = []class{
		class{
			[]string{"coding", "3min"},
			0,
		},
		class{
			[]string{"coding2", "3min2"},
			1,
		},
	}
}

func main() {
	firstPage := tables[0]
	err := firstPage.forEach(func(names []string) (bool, error) {
		if len(names) == 0 {
			return false, errors.New("")
		}
		//迭代每个数值
		for _, name := range names {
			fmt.Println(name)
		}
		return true, nil
	})

	if err != nil {
		fmt.Println(err)
	}

}
