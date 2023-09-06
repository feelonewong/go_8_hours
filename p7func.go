package main

import "fmt"

func main() {
	c := foo("abc", 123)
	fmt.Println("c:", c)

	// 返回多个参数
	d, e, f := bar("aaa", 123)
	fmt.Println("d=", d, ", e=", e, " ,f=", f)
	// 返回值带名称
	r4, r3 := backResponseParams()
	fmt.Println("r4:", r4, " ,r3:", r3)
}

func foo(a string, b int) int {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c := 100
	return c
}

// 多个匿名返回值
func bar(a string, b int) (int, string, bool) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c := 100
	d := "hello"
	flag := true
	return c, d, flag
}

// 多个具名返回值
func backResponseParams() (r1 int, r2 string) {
	// r1 r2有默认值 作用域空间是整个函数的作用域空间，默认值是0
	r1 = 1000
	r2 = "back to params"
	return r1, r2
}
