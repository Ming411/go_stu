package main

import "fmt"

func main() {
	// map结构   key value 的无序集合，主要用于查询
	var courseMap = map[string]string{
		"course1": "golang",
		"course2": "java",
		"course3": "python",
	}
	fmt.Println(courseMap["course1"]) // golang

	// 如果想后续添加值，必须初始化map给定初始值{}
	// 但是切片是可以直接进行append操作
	/*
		var courseMap = map[string]string{}
		var courseMap = make(map[string]string,3) // 常用
	*/
	courseMap["name"] = "coder"
	fmt.Println("courseMap:", courseMap)

	/* 遍历 map */
	for key, value := range courseMap {
		fmt.Println(key, value)
	}
	// 每次输出顺序不一定相同
	for key := range courseMap {
		fmt.Println(courseMap[key])
	}

	/* 	map 相关方法 */
	// 1. 判断对应的key是否存在map中，d值，ok是否存在
	// d, ok := courseMap["course2"]
	// fmt.Println(d, ok)

	if d, ok := courseMap["course2"]; !ok {
		fmt.Println("not in")
	} else {
		fmt.Println("in---", d)
	}
	// 2. 删除一个元素,即使元素不存在也不会报错
	delete(courseMap, "course2")
}
