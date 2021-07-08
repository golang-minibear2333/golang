// 切片陷阱
package main

import (
	"bytes"
	"fmt"
)

func main() {
	buffer := bytes.NewBuffer(make([]byte, 0, 100))
	buffer.Write([]byte("abc"))
	resBytes := buffer.Bytes()
	fmt.Printf("%s \n", resBytes)
	resBytes[0] = 'd'
	fmt.Printf("%s \n", resBytes)
	fmt.Printf("%s \n", buffer.Bytes())
}
