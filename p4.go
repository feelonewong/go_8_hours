package main

import (
	"fmt"
	"time"
)

func main() {
	// go build hello.go 编译生成.exe文件
	// go run hello.go  编译+运行
	fmt.Println("Hello World!")

	time.Sleep(1 * time.Second)
	fmt.Println(" Hello")
}
