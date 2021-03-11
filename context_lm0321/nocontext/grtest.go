package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	work(done)
	work(done)
	work(done)

	time.Sleep(3 * time.Second)
	close(done)
	time.Sleep(3 * time.Second)

}

func work(done chan interface{}) {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d\n", i)
			select {
			case <-done:
				fmt.Printf("OK, I quit!\n")
				return
			case <-time.After(time.Second):
			}
		}
	}()
}
