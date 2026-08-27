package main

import (
	"errors"
	"fmt"
	"strconv"
)

func validateAge(s string) (int, error) {
	if n, err := strconv.Atoi(s); err != nil {
		return 0, errors.New("parse: strconv.Atoi: parsing \"" + s + "\": invalid syntax")
	} else if n < 0 {
		return 0, errors.New("negative")
	} else {
		return n, nil
	}
}

func main() {
	var s string
	fmt.Scan(&s)
	n, err := validateAge(s)
	if err == nil {
		fmt.Println("age:", n)
	} else {
		fmt.Println("error:", err)
	}
}
