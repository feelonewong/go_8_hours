package main

import (
	"go_8_hours/pkg"
)

func main() {
	// 导包默认先执行包里面的init函数执行结束以后才会去执行新的业务函数
	pkg.PkgTest1()
	pkg.PkgTest2()

}
