package main

import "fmt"

func main() {
	// map的输出是无序的。
	// 声明方式1:
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Printf("myMap1是一个空map \n")
	}
	myMap1 = make(map[string]string, 3)
	myMap1["one"] = "Java"
	myMap1["two"] = "C++"
	myMap1["three"] = "golang"
	fmt.Printf("myMap1: %v\n", myMap1)
	// 声明方式2:
	myMap2 := make(map[string]string)
	myMap2["one"] = "Golang"
	myMap2["two"] = "JavaScript"
	myMap2["three"] = "TypeScript"
	// 声明方式3:
	myMap3 := map[string]string{
		"one":   "APP",
		"two":   "PC",
		"three": "Mini Program",
	}
	fmt.Printf("myApp3: %v\n", myMap3)
}
