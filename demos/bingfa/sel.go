package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.WaitGroup 是 Go 语言提供的一种同步原语，用于等待一组 goroutine 完成执行。
var wg sync.WaitGroup

func cpuInfo(stop chan struct{}) {
	defer wg.Done() //在当前函数结束时调用
	//  defer wg.Done() 来通知 sync.WaitGroup 当前 goroutine 执行完成
	// 当所有 goroutine 执行完毕后
	// sync.WaitGroup 就会解除阻塞，主 goroutine 可以继续执行下一步操作。
	for {
		select {
		case <-stop: // 当stop有值就退出
			fmt.Println("退出cpu监控")
			return
		default:
			time.Sleep(time.Second)
			println("cpu info")
		}
	}
}

func main() {
	var stop = make(chan struct{})

	wg.Add(1)
	go cpuInfo(stop)
	time.Sleep(time.Second * 3)
	stop <- struct{}{}
	wg.Wait() // 会等待上面的goroutine执行完毕 才会执行后续代码
}
