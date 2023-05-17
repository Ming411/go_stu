package main

import (
	"fmt"
	"sync"
)

var sum int32
var wg sync.WaitGroup
var lock sync.Mutex

// 锁 不能被复制  复制后就失去了锁的效果
func add() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		/*
			1.加载sum
			2.计算sum
			3.写入sum-->抢占 例如：-1先写入1后写入结果就是1  顺序不可控
		*/
		// fmt.Println("add:", sum)
		// sum += 1 // 可能出现抢占资源的情况
		/* 方式一 使用互斥锁 */
		lock.Lock()
		// fmt.Println("add:", sum)
		sum += 1
		lock.Unlock()
		/* 方式二 使用 atomic 原子化*/
		// atomic.AddInt32(&sum, 1)

	}
}
func sub() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		// fmt.Println("sub:", sum)
		// sum -= 1
		lock.Lock()
		// fmt.Println("sub:", sum)
		sum -= 1
		lock.Unlock()
		// atomic.AddInt32(&sum, -1)
	}
}

func main() {
	wg.Add(2) // 添加异步进程个数
	go add()
	go sub()
	wg.Wait()
	fmt.Println(sum)
}
