package tools

import "fmt"

func PrintSlice(arr []float32) {
	for _, v := range arr {
		fmt.Print(v)
	}
	fmt.Println()
}

func PrintMulti2Slice(arr [][]int) {
	// 注意遍历不定长切片，不可以用len来测长度，而要用cap，为了防止忘记建议用range最安全
	// 唯一的缺点多一个赋值的过程，浪费少量空间和时间（这是值得的）
	//for i := 0; i < len(arr); i++ {
	//	for j := 0; j < len(arr[j]); j++ {
	//		fmt.Printf("%v ", arr[i][j])
	//	}
	//	fmt.Println()
	//}

	for _, v1 := range arr {
		for _, v2 := range v1 {
			fmt.Printf("%v ", v2)
		}
		fmt.Println()
	}
}
