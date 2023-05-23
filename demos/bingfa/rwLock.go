package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var rwlock sync.RWMutex
	var wg sync.WaitGroup
	wg.Add(6)
	go func() {
		time.Sleep(time.Second * 3)
		defer wg.Done()
		rwlock.Lock() // 写入锁 会禁止其他写锁获取和读取
		defer rwlock.Unlock()
		fmt.Println("write+++")
		time.Sleep(time.Second * 5)
	}()
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for {
				rwlock.RLock()                     // 读锁   该锁并不会阻止别人读
				time.Sleep(500 * time.Millisecond) // 500ms
				fmt.Println("read---")
				rwlock.RUnlock()
			}
		}()
	}
	wg.Wait() // 等待 goroutine 执行完毕
}
