package main

import (
	"context"
	"log"
	"time"
)

func waitUntilTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	go func() { log.Println("Other goroutine") }()
	log.Println("Stop")
	<-ctx.Done()
	log.Println("Done")

	// 2022/12/28 00:16:21 Stop
	// 2022/12/28 00:16:21 Other goroutine
	// タイムアウトまで待つ
	// 2022/12/28 00:16:31 Done
}

func retryUntilCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	task := make(chan int)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case i := <-task:
				log.Println("get", i)
			default:
				log.Println("Not canceled yet")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()
	time.Sleep(1 * time.Second)
	for i := 0; i < 5; i++ {
		task <- i
	}
	cancel()

	// 2022/12/28 00:22:47 Not canceled yet
	// 2022/12/28 00:22:47 Not canceled yet
	// 2022/12/28 00:22:48 Not canceled yet
	// 2022/12/28 00:22:48 Not canceled yet
	// 2022/12/28 00:22:48 get 0
	// 2022/12/28 00:22:48 Not canceled yet
	// 2022/12/28 00:22:49 get 1
	// 2022/12/28 00:22:49 Not canceled yet
	// 2022/12/28 00:22:49 get 2
	// 2022/12/28 00:22:49 Not canceled yet
	// 2022/12/28 00:22:49 get 3
	// 2022/12/28 00:22:49 Not canceled yet
	// 2022/12/28 00:22:50 get 4
	// 2022/12/28 00:22:50 Not canceled yet
}

func main() {
	waitUntilTimeout()
	retryUntilCancel()
}
