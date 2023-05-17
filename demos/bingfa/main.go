package main

import "fmt"

func asyncPrint() {
	fmt.Println("async")
}

// 协程   goroutine
func main() {
	go asyncPrint() // 异步调用
	// 注意：主函数执行完毕，异步可能还没执行就随主程序一块销毁了
	fmt.Println("sync")

	for i := 0; i < 100; i++ {
		// go func() {
		// 	// 先生成的 并不一定就会先执行
		// 	fmt.Println(i) // 因为 goroutine是异步调用,所以输出是乱的,并且可能重复
		// }()

		/* 方式一 */
		// tmp := i // 可以避免重复，但是顺序依旧是乱的
		// go func() {
		// 	fmt.Println(tmp)
		// }()

		/* 方式二 */
		go func(i int) {
			fmt.Println(i) // 也可以避免重复，顺序随机
		}(i)
	}
}
