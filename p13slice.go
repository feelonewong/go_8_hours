package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice1 type: %T, value: %v\n", slice1, slice1)
	// 声明
	slice2 := make([]int, 3) // 默认开辟3个空间
	slice2[0] = 100
	fmt.Printf("slice2: %v\n", slice2)

	var slice3 []int
	if slice3 == nil {
		fmt.Println("slice3 是一个空切片")
	} else {
		fmt.Println("slice3 不是一个空切片")
	}
}
