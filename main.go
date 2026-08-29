package main

import (
	"fmt"
	"strconv"
)

type Logger struct {
}

func (l Logger) Log(msg string) {
	fmt.Println("[log]", msg)
}

type Counter struct {
	Logger
	count int
}

func (c *Counter) Inc() {
	c.Log(strconv.Itoa(c.count))
	c.count = c.count + 1
}

func main() {
	n := 0
	fmt.Scan(&n)
	c := Counter{Logger: Logger{}, count: 1}
	for i := 0; i < n; i++ {
		c.Inc()
	}
}
