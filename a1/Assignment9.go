package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 在函数结束时调用 Done，表示该工作已完成
	fmt.Printf("Worker %d is starting\n", id)
	time.Sleep(time.Second) // 模拟工作
	fmt.Printf("Worker %d is done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)         // 增加等待组计数
		go worker(i, &wg) // 启动 goroutine
	}

	wg.Wait() // 等待所有 goroutine 完成
	fmt.Println("All workers are done.")
}
