package main

import (
	"fmt"
	"sync"
)

var count int = 0

func main() {
	// 初始化锁
	lk := sync.Mutex{}

	done := make(chan bool)

	for i:=0; i < 100; i++ {
		// 并发的累加count
		go func() {
			// 加锁
			lk.Lock()
			// 延迟解锁
			defer lk.Unlock()

			// 处理业务逻辑
			count++
			done <- true
		}()
	}

	for i:=0; i < 100; i++ {
		<-done
	}
	fmt.Println(count)
}