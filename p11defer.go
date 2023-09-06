package main

import "fmt"

func main() {
	// 多个defer 顺序从后往前直行
	//defer fmt.Println("defer run ...")
	//defer fmt.Println("defer run 2...")
	//fmt.Println("main run 1...")
	//fmt.Println("main run 2...")
	deferAndReturn()
	// defer和return 同时存在的时候 defer最后执行，return在defer之前执行
}

func deferAndReturn() int {
	defer deferFunc()
	return returnFunc()
}
func deferFunc() int {
	fmt.Println("defer called ...")
	return 0
}

func returnFunc() int {
	fmt.Println(" return called ...")
	return 0
}
