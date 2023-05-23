package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cpuInfo(ctx context.Context) {
	defer wg.Done() //在当前函数结束时调用
	for {
		select {
		case <-ctx.Done(): // 当stop有值就退出
			fmt.Println("退出cpu监控")
			return
		default:
			time.Sleep(time.Second)
			println("cpu info")
		}
	}
}

func main() {
	wg.Add(1)
	//context包提供了三种函数，
	/* withCancel,
	WithTimeout,  // 主动超时
	withDeadline, // 在时间点cancel
	WithValue //
	*/
	ctx1, cancel := context.WithCancel(context.Background())
	ctx2, _ := context.WithCancel(ctx1)
	go cpuInfo(ctx2) // 具有传递性，父级cancel子集也会取消
	time.Sleep(time.Second * 3)
	cancel()
	wg.Wait()
}
