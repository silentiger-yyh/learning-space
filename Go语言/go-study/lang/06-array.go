package main

import "fmt"

func main() {
	var arr [3]int
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i]) // 未赋值，默认为0
	}

	arr1 := [3]int{1, 2, 3} // 不赋值会报错
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	arr2 := [2][2]int{{1, 2}, {3, 4}}
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[0]); j++ {
			fmt.Println(arr2[i][j])
		}
	}
}
