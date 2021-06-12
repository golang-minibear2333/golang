/*
* @Title:   灵活解析多版本字段类型稍微不同的json
* @Author:  minibear2333
* @Date:    2020-04-03 19:16
* @url:     https://github.com/golang-minibear2333/golang
*/
package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type People struct {
	Name string `json:"name"`
	Age  int    `json:"agent"`
}

//json 后面跟下划线，代表暂不接收这个字段值，后面的UnmarshalJSON方法来接收然后再赋值上去
type Coding3min struct {
	ID     string `json:"_"`
	People People `json:"_"`
	Total  int    `json:"total"`
}

/*
	重新实现Coding3min这个结构体的UnmarshalJSON方法，在json.Unmarshal时会被调用
	重写有两个作用：
		1. 做单位转换
		2. 兼容多个微版本同一个字段不同返回值
 */
func (r *Coding3min) UnmarshalJSON(b []byte) error {
	//用来接受非`json:"_"`的字段
	type tmp Coding3min
	/*
		用中间变量接收json串，tmp以外的字段用来接受`json:"_"`属性字段
		注意必须为interface{}类型，这样才可以接收任意字段
	 */
	var s struct {
		tmp
		ID     interface{} `json:"id"`
		People interface{} `json:"people"`
	}

	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	*r = Coding3min(s.tmp)

	var tmpb []byte
	/*
		people可能有两种值：
		1. "{\"name\":\"tom\",\"age\":12}" 此时他是字符串类型
		2. {"name":"tom","age":12} 	此时他就是struct类型
	 */
	switch t := s.People.(type) {
	case string:
		tmpb = []byte(t)
		//这里就要打开json.Unmarshal的源码看了，struct解析以后是这种类型
	case map[string]interface{}:
		tmpb, err = json.Marshal(t)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("People has unexpected type: %T", t)
	}

	if len(tmpb) != 0 {
		err = json.Unmarshal(tmpb, &r.People)
		if err != nil {
			return err
		}
	}

	/*
		id可以是三种类型的值：int float64 string
	 */
	switch t := s.ID.(type) {
	case int:
		r.ID = strconv.Itoa(t)
	case float64:
		r.ID = strconv.Itoa(int(t))
	case string:
		r.ID = t
	default:
		return fmt.Errorf("ID has unexpected type: %T", t)
	}

	return nil
}
func parseJson(jsonStr string) {
	//这一步是有必要的，先把json类型先解析成interface类型中的隐式struct类型
	var tmp interface{}
	err := json.Unmarshal([]byte(jsonStr), &tmp)
	if err != nil {
		fmt.Println("parseJson Unmarshal error:" + err.Error())
		return
	}
	//把interface类型转换成我们想要的struct类型
	var coding3min Coding3min
	err = ExtractInto(tmp, &coding3min)
	if err != nil {
		fmt.Println("parseJson ExtractInto error:" + err.Error())
		return
	}
	printStr := fmt.Sprintf("%v", coding3min)
	fmt.Println(printStr)
}

/*
	把interface类型转换成我们想要的struct类型
	这个通用方法可以转换成任意一个想要的类型
 */
func ExtractInto(source interface{}, to interface{}) error {

	b, err := json.Marshal(source)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, to)

	return err
}

func main() {
	jsonStr := `{"id":12,"people":"{\"name\":\"tom\",\"age\":12}","total":1}`
	parseJson(jsonStr)
	jsonStr = `{"id":"12","people":{"name":"tom","age":12},"total":1}`
	parseJson(jsonStr)
}
