package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 0
	fmt.Scan(&n)
	num := []int{}
	for i := 0; i < n; i++ {
		t := 0
		fmt.Scan(&t)
		num = append(num, t)
	}
	each := n / 4
	total := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for c := 0; c < each; c++ {
			mu.Lock()
			total += num[c]
			mu.Unlock()
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for c := each; c < each*2; c++ {
			mu.Lock()
			total += num[c]
			mu.Unlock()
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for c := each * 2; c < each*3; c++ {
			mu.Lock()
			total += num[c]
			mu.Unlock()
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for c := each * 3; c < n; c++ {
			mu.Lock()
			total += num[c]
			mu.Unlock()
		}
	}()
	wg.Wait()
	fmt.Println(total)
}
