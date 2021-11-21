package lib2

import "fmt"

func Lib1Test() {
	fmt.Println("lib1Test()...")
}

func init() {
	fmt.Println("lib2. init()...")
}