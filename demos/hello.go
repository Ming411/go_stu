package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println("hello world")

	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum=", sum)

	// 99乘法表
	for i := 1; i < 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println("")
	}

	// var name string = "coder"

	var strtoint string = "123"
	// 字符串转数字，可能出错，所以要捕获
	ret, err := strconv.Atoi(strtoint)
	if err != nil {
		return

	}
	fmt.Println(ret)

	// 格式化输出

	nickname := "coder"
	age := 18
	gender := "male"
	fmt.Printf("nickname=%s,age=%d,gender=%s", nickname, age, gender)

	var arr [3]int = [3]int{
		1, 2, 3,
	}

	fmt.Println(arr)
}
