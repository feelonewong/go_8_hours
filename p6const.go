package main

func main() {
	// 常量 1.不可修改  2.const关键字定义枚举类型
	const LENGTH int = 10

	// iota关键字 每行都会累加1
	const (
		BEIJING = iota
		SHANGHAI
		GUANGZHOU
		SHENZHEN
	)
	// iota 只能配合const使用
	const (
		NANJING = 10 * iota
		HEFEI
		CHENGDU
		XIAN
	)
}
