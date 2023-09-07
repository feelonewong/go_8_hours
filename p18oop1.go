package main

import "fmt"

type Hero struct {
	Name string
}

/* 1.结构体属性只有是大写 才是公开属性
 * 2. 结构体方法最好是使用指针 这样才是引用传递
 */
func (this Hero) GetName() string {
	return this.Name
}
func (this *Hero) SetName(name string) {
	this.Name = name
}

func main() {
	hero := Hero{
		Name: "Huang LongFei",
	}
	fmt.Println("Hero:", hero)
	hero.SetName("黄飞龙")
	fmt.Println(hero.GetName())

	hero2 := Hero{
		Name: "小黄",
	}
	fmt.Println(hero2.GetName())
}
