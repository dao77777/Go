package main

import ("fmt")

var gA int = 100
var gB = 200

const (
	BEIJING = iota
	SHANGHAI
	SHENZHEN
)

const (
	WOW, HAHA = iota + 1, iota + 2
	WUHU, YOYO
)

func main() {
	var a int
	fmt.Println("a =", a)
	fmt.Printf("a = %T\n", a)
	var b int = 100
	fmt.Println("b =", b)
	fmt.Printf("type of b = %T\n", b)
	var c = 200
	fmt.Println("c =", c)
	fmt.Printf("type of c = %T\n", c)
	var d string = "dao77777"
	fmt.Println("d =", d)
	fmt.Printf("d = %s, type of d = %T", d, d)
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)
	g := 3.14
	fmt.Println("g = ", g)
	fmt.Printf("type of g = %T\n", g)
	const f int = 100
	fmt.Printf("f = %d, type of f = %T", f, f)
	var xx, yy int = 100, 200
	fmt.Println("xx =", xx, "yy =", yy)
	var kk, ll = 100, "haha"
	fmt.Println("kk =", kk, "ll =", ll)
	var (
		vv int = 100
		jj bool = true
	)
	fmt.Println("vv =", vv, "jj =", jj)

	fmt.Println("gA =", gA, ",gB =", gB)
	fmt.Println("BEIJING =", BEIJING, "SHENZHEN =", SHENZHEN, "SHANGHAI =", SHANGHAI)
	fmt.Println("WOW =", WOW, "HAHA =", HAHA, "WUHU =", WUHU, "YOYO =", YOYO)
}