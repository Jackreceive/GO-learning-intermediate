package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func cal(a int, b int) (x int, err error) {
	defer func() {
		if t := recover(); t != nil {
			err = fmt.Errorf("divide by zero")
		}
	}()
	return a / b, nil
}

func main() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	m, _ := strconv.Atoi(r.Text())
	r.Scan()
	n, _ := strconv.Atoi(r.Text())
	if ans, err := cal(m, n); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", ans)
	}
}
