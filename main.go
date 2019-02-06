package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancelFunc := context.WithTimeout(ctx, time.Second) // increase time to see difference
	defer cancelFunc()

	ch := make(chan bool)
	go wait(ch)

	fmt.Println("work run select")
	select {
	case <-ch:
		fmt.Println("Wait() done") // wait() finished before timeout
	case <-ctx.Done():
		fmt.Println("Context done") // timeout
	}
}

func wait(ch chan bool) {
	fmt.Println("wait starts")
	time.Sleep(2 * time.Second)
	fmt.Println("wait write in channel")
	ch <- true
	fmt.Println("wait done")
}
