package main

import (
	"fmt"
	"io"
	"os"
)

type Animal interface {
	Sleep()
	GetColor() string
	GetType() string
}

// Cat
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// Dog
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

//
func PrintAnimal(animal Animal) {
	fmt.Printf("color = %s, type = %s\n", animal.GetColor(), animal.GetType())
}

func myFunc(arg interface{}) {
	val, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, val =", val)
	}
	fmt.Println("myFunc is called...")
	fmt.Println(arg)
}

func main() {
	var animal Animal = &Cat{"black"}
	fmt.Printf("type of animal = %T\n", animal)
	PrintAnimal(animal)
	myFunc("wow")

	tty, err := os.OpenFile("./test.md", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file fail", err)
	}
	fmt.Println(tty)
	fmt.Printf("%T\n", tty)

	var r io.Reader
	r = tty

	var w io.Writer
	w = r.(io.Writer)

	w.Write([]byte("Hello"))
}
