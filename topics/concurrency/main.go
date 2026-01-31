package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// A compact, runnable tutorial demonstrating common usages of
// Go concurrency primitives. Run with `go run main.go`.

func main() {
	fmt.Println("Go concurrency tutorial — quick examples")
	fmt.Println("-----------------------------------------")

	goroutinesWaitGroup()
	fmt.Println()

	channelsSelect()
	fmt.Println()

	mutexSharedState()
	fmt.Println()

	contextCancellation()
}

func goroutinesWaitGroup() {
	fmt.Println("1) Goroutines + WaitGroup")
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(id) * 50 * time.Millisecond)
			fmt.Println("goroutine", id, "done")
		}(i)
	}
	wg.Wait()
}

func channelsSelect() {
	fmt.Println("2) Channels and select (producer/consumer)")
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Println("recv", v)
		case <-time.After(200 * time.Millisecond):
			fmt.Println("timeout waiting for channel")
			return
		}
	}
}

func mutexSharedState() {
	fmt.Println("3) Mutex for shared state")
	var mu sync.Mutex
	counter := 0
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("counter:", counter)
}

func contextCancellation() {
	fmt.Println("4) Context for cancellation")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan string)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine cancelled")
				return
			default:
				time.Sleep(50 * time.Millisecond)
				ch <- "working"
			}
		}
	}()

	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
	cancel()
	time.Sleep(100 * time.Millisecond) // wait for cancellation
}