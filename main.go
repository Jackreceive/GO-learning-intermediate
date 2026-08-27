package main

import (
	"fmt"
	"sync"
)

var ans, N int
var mutex sync.Mutex
var num []int
var wg sync.WaitGroup

func goroutine(a int) {
	defer wg.Done()
	res := 0
	i := N / 4 * a
	for ; i < N/4*(a+1); i++ {
		res += num[i]
	}
	if a == 3 {
		for ; i < N; i++ {
			res += num[i]
		}
	}
	mutex.Lock()
	ans += res
	mutex.Unlock()

}

func main() {
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		t := 0
		fmt.Scan(&t)
		num = append(num, t)
	}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go goroutine(i)
	}
	wg.Wait()
	fmt.Println(ans)
}
