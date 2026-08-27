package main

import "fmt"

type shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Square struct {
	Side float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func main() {
	var s string
	var n float64
	fmt.Scan(&s, &n)
	if s == "square" {
		t := Square{n}
		fmt.Printf("%.2f", t.Area())
		return
	}
	t := Circle{n}
	fmt.Printf("%.2f", t.Area())

}
