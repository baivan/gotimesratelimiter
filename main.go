package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	stopper := false
	done := make(chan bool)

	go func() {
		//reset:
		i := 0
		for {
			if !stopper && i < 2 {
				time.Sleep(500 * time.Millisecond)
				sendMessage(i)
				i++
			} else {
				fmt.Printf("Reseting at %d\n", i)
				stopper = false
				i = 0
			}
		}
	}()

	go func() {
		for {
			select {
			case t := <-ticker.C:
				fmt.Printf("Ticked here %s\n", t)
				stopper = true
			}
		}
	}()

	<-done
}

func sendMessage(count int) {
	fmt.Printf("Performing counted operation %d\n", count)
}
