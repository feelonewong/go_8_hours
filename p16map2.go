package main

import "fmt"

func main() {
	// map是引用类型数据
	// 依次实现map的增删改查
	cityMap := make(map[string]string)
	cityMap["AnHui"] = "HeFei"
	cityMap["JiangSu"] = "NanJing"
	cityMap["ZheJiang"] = "HangZhou"

	// 增加
	updateMap(cityMap, "GuangDong", "GuangZhou")
	// 删除
	deleteMap(cityMap, "JiangSu")
	// 修改
	updateMap(cityMap, "AnHui", "NanJing")
	// 查找
	queryMap(cityMap)
}
func updateMap(newMap map[string]string, currentKey, currentValue string) {
	newMap[currentKey] = currentValue
}
func queryMap(newMap map[string]string) {
	for key, value := range newMap {
		fmt.Printf("newMap's key: %s, value: %s\n", key, value)
	}
}
func deleteMap(newMap map[string]string, key string) {
	delete(newMap, key)
}
