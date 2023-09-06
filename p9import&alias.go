package main

import myPkg "go_8_hours/pkg"

func main() {
	// 导入的包一定要使用
	// 匿名的导包方式 import _ "xxx" 可以直接不使用包下面的函数也不报错
	// 别名 import myLib "xx/lib"
	// 全局引入 import . "go_8_hours/pkg" 谨慎使用
	myPkg.PkgTest2()
}
