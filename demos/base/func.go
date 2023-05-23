package main

import (
	"fmt"
	"time"
)

/* go 语言中的函数 */

// 都是值传递，内部会进行一层拷贝操作
func add(a, b int) (int, error) {
	// 可以有多个返回值
	return a + b, nil
}
func add2(a, b int) (sum int, err error) {
	// 可以直接给返回值命名
	sum = a + b
	return sum, err
}
func add3(item ...int) (sum int, err error) {
	// fmt.Println(item) //  item 是一个切片
	for _, value := range item {
		sum += value
	}
	return // 默认返回sum 和 err
}
func runForever() {
	for {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
	// 也可以没有返回值
}

/* 闭包示例 */
func autoIncrement() func() int {
	local := 0
	return func() int {
		local += 1
		return local
	}
}

func main() {
	sum, _ := add(1, 2)
	fmt.Println(sum)

	sum2, _ := add3(1, 2, 3, 4, 5)
	fmt.Println(sum2) // 15

	nextFunc := autoIncrement()
	for i := 0; i < 3; i++ {
		fmt.Println(nextFunc()) // 1 2 3
	}
}
