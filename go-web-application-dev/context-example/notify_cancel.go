package main

import (
	"context"
	"log"
)

func child(ctx context.Context) {
	log.Println("Child started")
	if err := ctx.Err(); err != nil {
		return
	}
	log.Println("Not canceled")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	child(ctx)
	// 2022/12/28 00:09:22 Child started
	// 2022/12/28 00:09:22 Not canceled
	cancel()
	child(ctx)
	// 2022/12/28 00:09:22 Child started
}
