package main

import (
	"fmt"
)

type myint int

type Book struct {
	title string
	auth string
}

func printBook(book Book) {
	fmt.Printf("title = %s, auth = %s\n", book.title, book.auth)
}

func changeBook(book *Book) {
	book.title = "wo"
}

func main() {
	var a myint = 10
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	var book Book
	book.title = "GoLang"
	book.auth = "zhang3"
	printBook(book)
	changeBook(&book)
	printBook(book)
}