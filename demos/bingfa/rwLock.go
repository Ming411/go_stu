package main

import (
	"fmt"
	"sync"
	"time"
)

/* 读写锁 */
func main() {
	var rwLock sync.RWMutex
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		time.Sleep(3 * time.Second)
		rwLock.Lock() // 写锁    可以防止其他写锁获取和读锁获取
		fmt.Println("write++++")
		rwLock.Unlock()
		wg.Done() // 先压栈 后弹栈
		time.Sleep(time.Second * 5)
	}()

	go func() {
		defer wg.Done()
		for {
			rwLock.RLock() // 读锁  读锁不会组织别人读取
			time.Sleep(500 * time.Millisecond)
			fmt.Println("read----")
			rwLock.RUnlock()

		}
	}()

	wg.Wait()

}
