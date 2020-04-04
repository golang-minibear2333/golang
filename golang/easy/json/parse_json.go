/*
* @Title:   json解析与编码
* @Author:  pzqu
* @Date:    2020-03-15 11:46
* @url:     https://github.com/pzqu/how_to_code
*/

package main

import (
	"encoding/json"
	"fmt"
)

type Coding3min struct {
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
func parseJson(jsonStr string) (Coding3min, error) {
	var coding3min Coding3min
	err := json.Unmarshal([]byte(jsonStr), &coding3min)
	return coding3min, err
}

//把struct编译成string
func convertJson(c Coding3min) (res string, err error) {
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
