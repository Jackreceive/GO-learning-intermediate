package main

import (
	"fmt"
	"time"
)

func main() {
	a := make(chan string, 2)
	go func() {
		time.Sleep(30 * time.Millisecond)
		a <- "slow"
	}()
	go func() {
		time.Sleep(10 * time.Millisecond)
		a <- "fast"
	}()
	select {
	case t := <-a:
		fmt.Println(t)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timeout")
	}
}
