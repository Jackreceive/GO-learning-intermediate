package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	a := make(chan int)
	b := make(chan int)
	fields := strings.Fields(r.Text())
	var num []int
	for _, s := range fields {
		n, _ := strconv.Atoi(s)
		num = append(num, n)
	}
	go func() {
		defer close(a)
		for i := 0; i < len(num); i++ {
			a <- num[i]
		}
	}()
	go func() {
		defer close(b)
		for i := range a {
			b <- i * i
		}
	}()
	ans := 0
	for i := range b {
		ans += i
	}
	fmt.Println(ans)

}
