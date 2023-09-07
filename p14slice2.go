package main

import "fmt"

func main() {
	numbers := make([]int, 3, 5)
	fmt.Printf("len: %d, cap: %d, numbers: %v\n", len(numbers), cap(numbers), numbers)

	// 切片追加 len:4 cap:5
	numbers = append(numbers, 1)
	fmt.Printf("len: %d, cap: %d, numbers: %v\n", len(numbers), cap(numbers), numbers)
	// len: 5 cap:5
	numbers = append(numbers, 1)
	fmt.Printf("len: %d, cap: %d, numbers: %v\n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 1)
	// 追加的长度超出cap长度时，cap的长度是之前cap的2倍 len:6 cap:10
	fmt.Printf("len: %d, cap: %d, numbers: %v\n", len(numbers), cap(numbers), numbers)

	// 切片截取:
	s1 := numbers[0:2] // 左包括，右不包括
	fmt.Println("s1:", s1)
	// copy解决引用问题 深拷贝
	c1 := []int{1, 2, 3, 4, 5}
	c2 := make([]int, 5)
	copy(c2, c1)
	c2[0] = 100
	fmt.Printf("c1: %v, c2: %v \n", c1, c2)
}
