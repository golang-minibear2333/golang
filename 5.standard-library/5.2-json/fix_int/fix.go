package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type People struct {
	Name string `json:"name"`
	Age  int    `json:"_"`
}

func (p *People) UnmarshalJSON(b []byte) error {
	// 定义临时类型 用来接受非`json:"_"`的字段
	type tmp People
	// 用中间变量接收json串，tmp以外的字段用来接受`json:"_"`属性字段
	var s = &struct {
		tmp
		// interface{}类型，这样才可以接收任意字段
		Age interface{} `json:"age"`
	}{}
	// 解析
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	// 判断真实类型，并类型转换
	switch t := s.Age.(type) {
	case string:
		var age int
		age, err = strconv.Atoi(t)
		if err != nil {
			return err
		}
		s.tmp.Age = age
	case float64:
		s.tmp.Age = int(t)
	}
	// tmp类型转换回People，并赋值
	*p = People(s.tmp)
	return nil
}

func main() {
	var req1Json = []byte(`{"age":26}`)
	var req2Json = []byte(`{"age":"26"}`)
	var person People
	err := json.Unmarshal(req1Json, &person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(person)

	err = json.Unmarshal(req2Json, &person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(person)
}
