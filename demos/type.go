package main

import "fmt"

func main() {
	// 1.类型别名
	type Myint = int
	var a Myint           // 在编译的时候会自动编译为 int
	fmt.Printf("%T\n", a) // int

	// 2.自定义类型
	// type Myint int  // 可以扩展int类型的方法

	var b interface{} = "hello"
	switch b.(type) {
	case string:
		fmt.Println("字符串") //
	}
}
