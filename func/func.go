package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	c := 100

	return c
}

func foo2(a string, b int) (int, int) {
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	return 666, 777
}

func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("---- fool3 ----")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	r1 = 1000
	r2 = 2000

	return
}

func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("---- fool3 ----")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	r1 = 1000
	r2 = 2000

	return
}

func main() {
	c := foo1("abc", 555)
	fmt.Println("c =", c)

	ret1, ret2 := foo2("haha", 999)
	fmt.Println("ret1 =", ret1, "ret2 =", ret2)

	ret1, ret2 = foo3("foo3", 333)
	fmt.Println("ret1 =", ret1, "ret2 =", ret2)
}