package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1000 * time.Millisecond)

	var quitFlag string
	fmt.Scanln(&quitFlag)
	// if quitFlag == "q" {
	ticker.Stop()
	done <- true
	// }
	fmt.Println("Ticker stopped")
}
