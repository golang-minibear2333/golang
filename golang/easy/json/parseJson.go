/*
* @Author: pzqu
* @Date:   2020-03-15 11:46
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

func parseJson(jsonStr string) (Coding3min, error) {
	var coding3min Coding3min
	err := json.Unmarshal([]byte(jsonStr), &coding3min)
	return coding3min, err
}

func main() {
	jsonStr := `{"data":{"items":[{"_id":"2"}],"total_count":1},"message":"","result_code":200}`
	coding3min, err := parseJson(jsonStr)
	if err != nil {
		fmt.Println("parseJsonErr error:" + err.Error())
		return
	}
	printStr := fmt.Sprintf("%v", coding3min)
	fmt.Println(printStr)
}
