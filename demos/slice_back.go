package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5}
	s1 := data[0:3]
	s2 := data[1:4] // 当没有出现被动扩容情况下，引用调用

	// cap容量，从原切片开始位置到结束即为容量
	fmt.Println(len(s1), cap(s1)) // 3 5
	fmt.Println(len(s2), cap(s2)) //3 4

	// s2 = append(s2, 6, 7, 8, 9, 10) // 此时扩容后的切片为一个新的引用
	s2[0] = 200
	fmt.Println("s1:", s1)
}
