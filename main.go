package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	num []int
}

func (s *Stack) push(a int) {
	s.num = append(s.num, a)
}

func (s *Stack) pop() (int, bool) {
	if len(s.num) <= 0 {
		return 0, false
	}
	x := s.num[len(s.num)-1]
	s.num = s.num[:len(s.num)-1]
	return x, true
}

func main() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	field := strings.Fields(r.Text())
	var s Stack
	for _, v := range field {
		t, _ := strconv.Atoi(v)
		s.push(t)
	}
	for {
		n, b := s.pop()
		if b == false {
			break
		}
		fmt.Println(n)
	}

}
