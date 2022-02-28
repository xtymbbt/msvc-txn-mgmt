package test

import "fmt"

func Slicetest() {
	a := []int{1, 2, 3}
	a = append(a[:1], a[2:]...)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("len(a): %v\n", len(a))
	b := 5
	a = append([]int{5}, a...)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("len(a): %v\n", len(a))
	var x = make([]int, 0, 0)
	x = append(x, a[:2]...)
	fmt.Printf("x: %v\n", x)
	fmt.Printf("len(x): %v\n", len(x))
	x = append(x, b)
	y := a[2:]
	fmt.Printf("x: %v\n", x)
	fmt.Printf("len(x): %v\n", len(x))
	fmt.Printf("y: %v\n", y)
	fmt.Printf("len(y): %v\n", len(y))
	a = append(x, y...)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("len(a): %v\n", len(a))
}
