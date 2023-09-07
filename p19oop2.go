package main

import "fmt"

// 定义一个基类
type Animal struct {
	name string
	age  int
}

// 定义基类的方法
func (a *Animal) Eat() {
	fmt.Println("The animal is eating.")
}

// 定义一个派生类，组合基类
type Dog struct {
	animal Animal
	breed  string
}

// 重写派生类的方法
func (d *Dog) Eat() {
	fmt.Println("The dog is eating.")
}

// 添加派生类的新方法
func (d *Dog) Bark() {
	fmt.Println("The dog is barking.")
}

func main() {
	// 创建一个派生类对象
	dog := Dog{
		animal: Animal{
			name: "Bobby",
			age:  3,
		},
		breed: "Golden Retriever",
	}

	// 调用派生类的方法
	dog.Eat()  // 输出: The dog is eating.
	dog.Bark() // 输出: The dog is barking.

	// 调用基类的方法
	dog.animal.Eat() // 输出: The animal is eating.
}
