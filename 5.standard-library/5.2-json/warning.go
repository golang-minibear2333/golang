package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string `json:"name"`
	age  int    `json:"age"`
}

func err1() {
	reqJson := `{"name":"minibear2333","age":26}`
	var person People
	err := json.Unmarshal([]byte(reqJson), &person)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(person)
}

func err2() {
	raw := []byte(`{"name":"\\xc2"}`)
	var person People
	if err := json.Unmarshal(raw, &person); err != nil {
		fmt.Println(err)
	}
	fmt.Println(raw)
}

func err4() {
	var data = []byte(`{"age": 26}`)

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("error:", err)
		return
	}

	var status = result["age"].(int) //error
	fmt.Println("status value:", status)
}

func main() {
	err1()
	err2()
	err4()
}
