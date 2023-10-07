package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	go func() {
		ch <- 1
	}()

	select {
	case <-ch:
		fmt.Println("Normal")
	case <-time.After(300 * time.Millisecond):
		fmt.Println("Should be buggy")
	}
}
