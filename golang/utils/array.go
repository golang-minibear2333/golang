/*
* @Author: pzqu
* @Date:   2020-04-03 10:52
*/
package utils

import "fmt"

func PrintArray(arr []float32){
	for _,v :=  range arr{
		fmt.Print(v)
	}
	fmt.Println()
}
