package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker("mine worker", &wg)

	go worker("wine worker", &wg)

	wg.Wait()

	fmt.Printf("main exit ...")
}

func worker(name string, wg *sync.WaitGroup)  {
	time.Sleep(1 * time.Second)
	fmt.Printf("starting to work --- %s\n", name)
	wg.Done()

}