package main

import "fmt"

type IAnimal interface {
	Sleep()
	GetColor()
}

type Cat struct {
	Name  string
	Color string
}

func (this *Cat) Sleep() {
	fmt.Println(this.Name, ",睡觉了...")
}
func (this *Cat) GetColor() {
	fmt.Println(this.Name, "的颜色是:", this.Color)

}

func main() {
	// 自动继承interface：只要实现interface的方法就可以实现这个效果
	// 多态： 这里实现了接口，同一种行为多种形态 简称为多态
	cat := Cat{"小猫咪", "三花色"}
	cat.Sleep()
	cat.GetColor()

	var animal IAnimal
	animal = &Cat{"大橘子", "黄色"}
	animal.Sleep()
	animal.GetColor()
}
