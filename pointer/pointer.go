package main

import (
	"fmt"
)

func swap(pa *int, pb *int) {
	var t int = *pa
	*pa = *pb
	*pb = t
	return
}

func main() {
	var a int = 10
	var b int = 20

	swap(&a, &b)

	fmt.Println("a =", a, "b =", b)

	var p = &a
	fmt.Println("p =", p, "&a =", &a)
	fmt.Printf("type of p = %T\n", p)
	var pp = &p
	fmt.Println("pp =", pp)
	fmt.Printf("type of pp = %T\n", pp)
}