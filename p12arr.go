package main

import "fmt"

func main() {
	// 固定长度
	// 数组是值传递的方式，是值拷贝的形式。
	var myArr1 [10]int
	fmt.Println(myArr1)
	// 不定长度,动态数组
	// 切片:引用传递 与数组区分：不定长度、是引用类型数据。
	myArr2 := []int{10, 20, 30, 40, 50}
	fmt.Printf("myArr2 type  is: %T \n", myArr2)
	fmt.Println(myArr2)
	rangeArr(myArr2)
	fmt.Println(myArr2)
}

// 数组循环的方式1：
func rangeArr(arr []int) {
	for idx, value := range arr {
		value += 20
		fmt.Println("idx: ", idx, " , value: ", value, "\n")
	}
}

// 数组循环的方式2:
func forArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println("i: ", i, " , value: ", arr[i], "\n")
	}
}
