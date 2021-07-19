/*
* @Title:   json解析与编码
* @Author:  minibear2333
* @Date:    2020-03-15 11:46
* @url:     https://github.com/golang-minibear2333/golang
 */

package main

import (
	"encoding/json"
	"fmt"
)

// 注意初学者最容易犯的两个致命错误，解析不出来某个字段，还不会报错
// 1. 注意！！！ 要被解析python的结构体，成员变量名称必须首字母大写（权限问题），否则会因为访问不到无法解析成功
// 2. 注意！！！ 不要定义两个 类似 `json:"_id"` 相同却不同名的字段，不然找到死也找不到为什么解析不出来
type MiniBear2333 struct {
	Data struct {
		Items []struct {
			ID string `json:"_id"`
		} `json:"items"`
		TotalCount int64 `json:"total_count"`
	} `json:"data"`
	Message    string `json:"message"`
	ResultCode int64  `json:"result_code"`
}

//把string解析成struct
func (m MiniBear2333) parseJson(jsonStr string) (MiniBear2333, error) {
	var coding3min MiniBear2333
	err := json.Unmarshal([]byte(jsonStr), &coding3min)
	return coding3min, err
}

//把struct编译成string
func convertJson(c MiniBear2333) (res string, err error) {
	resBytes, err := json.Marshal(c)
	return string(resBytes), err
}

func main() {
	jsonStr := `{"data":{"items":[{"_id":"2"}],"total_count":1},"message":"","result_code":200}`
	//解析
	coding3min, err := parseJson(jsonStr)
	if err != nil {
		fmt.Println("parseJson error:" + err.Error())
		return
	}
	printStr := fmt.Sprintf("%v", coding3min)
	fmt.Println(printStr)
	//编译
	resStr, err := convertJson(coding3min)
	if err != nil {
		fmt.Println("convertJson error: " + err.Error())
	}
	fmt.Println(resStr)
}
