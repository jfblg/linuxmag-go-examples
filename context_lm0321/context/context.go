package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	work(ctx)
	work(ctx)
	work(ctx)

	time.Sleep(3 * time.Second)
	cancel() // from context package
	time.Sleep(3 * time.Second)

}

func work(ctx context.Context) {

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d\n", i)
			select {
			case <-ctx.Done():
				fmt.Printf("OK, I quit!\n")
				return
			case <-time.After(time.Second):
			}
		}
	}()

}
