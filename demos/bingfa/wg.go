package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(100) // 表示有100个需要等待的任务
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			defer wg.Done() // Add 和 Done  必须成对出现
		}(i)
	}

	wg.Wait() // 等待异步执行完毕，避免主程序结束而异步代码未执行
	fmt.Println("all done")
}
