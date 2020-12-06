package main

import (
	"fmt"
	"time"
)

const timeout = 1000

func main() {
	stop := make(chan bool)

	go worker("mine worker", stop)

	// insert true signal into the stop channel
	time.Sleep(time.Second * 5)
	stop <- true
	fmt.Printf("the work is done !\n")

}

func worker(name string, stop chan bool) {
	for true {
		select {
		case <-stop:
			fmt.Printf("end the work --- %s\n", name)
			return
		default:
			time.Sleep(time.Second)
			fmt.Printf("starting to work --- %s\n", name)

		}
	}
}
