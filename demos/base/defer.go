package main

import "fmt"

/* defer 语句只有在碰到return之前才会执行 */

func deferFunc() (ret int) {
	defer func() {
		ret++
	}()
	return 10
}

func main() {
	ret := deferFunc()
	// defer是有能力修改函数返回值的
	fmt.Println(ret) // 11
}

// func main() {
// 	defer fmt.Println("defer1")
// 	defer fmt.Println("defer2")
// 	fmt.Println("defer3")
// 	return
// 	//  输出  321
// 	// 会有一个压栈操作，弹栈是从外到里的
// }
