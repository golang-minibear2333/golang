package main

import (
	"encoding/json"
	"fmt"
)

// 注意初学者最容易犯的两个致命错误，解析不出来某个字段，还不会报错
// 1. 注意！！！ 要被解析python的结构体，成员变量名称必须首字母大写（权限问题），否则会因为访问不到无法解析成功
// 2. 注意！！！ 不要定义两个 类似 `json:"_id"` 相同却不同名的字段，不然找到死也找不到为什么解析不出来
type ResponseData struct {
	Data struct {
		Items      []Body `json:"items"`
		TotalCount int64  `json:"total_count"`
	} `json:"data"`
	Message    string `json:"message"`
	ResultCode int64  `json:"result_code"`
}

type Body struct {
	ID int `json:"_id"`
}

func foo() {
	jsonStr := `{"data":{"items":[{"_id":2}],"total_count":1},"message":"","result_code":200}`
	//把string解析成struct
	var responseData ResponseData
	err := json.Unmarshal([]byte(jsonStr), &responseData)
	if err != nil {
		fmt.Println("parseJson error:" + err.Error())
		return
	}
	fmt.Println(responseData)
}

func bar() {
	r := ResponseData{
		Data: struct {
			Items      []Body `json:"items"`
			TotalCount int64  `json:"total_count"`
		}{
			Items: []Body{
				{ID: 1},
				{ID: 2},
			},
			TotalCount: 1,
		},
		Message:    "",
		ResultCode: 200,
	}
	//把struct编译成string
	resBytes, err := json.Marshal(r)
	if err != nil {
		fmt.Println("convertJson error: " + err.Error())
	}
	fmt.Println(string(resBytes))
}

func fooErr() {
	jsonStr := `{"data":{"items":[{"_id":2}],"total_count":1},"message":"","result_code":200}`
	//把string解析成struct
	var responseData ResponseData
	err := json.Unmarshal([]byte(jsonStr), responseData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(responseData)
}
func main() {
	foo()
	bar()
	fooErr()
}
