package main

import "fmt"

func main() {
	a := 1
	changeValue(&a)
	fmt.Println(a)

	i := 10
	j := 20
	fmt.Println("i=", i, ",j=", j)
	swap(&i, &j)
	fmt.Println("i=", i, ",j=", j)
}
func changeValue(p *int) {
	*p = 10
}

// 交换值
func swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}
