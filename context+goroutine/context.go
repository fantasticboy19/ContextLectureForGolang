package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// we dont use the stop channel to propagate the signal to main goroutine
// we use root context to implement the functionality
func main() {
	// we use the cancel function to cancel goroutines which are belonged to the RootContext
	RootCtx, cancel := context.WithCancel(context.Background())

	go worker(RootCtx, "mine worker")
	go worker(RootCtx, "dinner worker")
	go worker(RootCtx, "lunch worker")
	go worker(RootCtx, "green tea worker")

	// and we use the cancel() to replace the stop channel
	time.Sleep(5 * time.Second)
	cancel()

	log.Printf("all mine workers are done, then we should exit the main go routine")
}


func worker(ctx context.Context, name string)  {
	for true {
		select {
		case <-ctx.Done():
			fmt.Printf("end the work --- %s\n", name)
			return
		default:
			time.Sleep(time.Second)
			fmt.Printf("starting to work --- %s\n", name)

		}
	}
}