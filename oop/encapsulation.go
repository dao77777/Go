package main

import (
	"fmt"
)

// 定义类
type Hero struct {
	name string
	ad int
	level int
}

func (this Hero) show() {
	fmt.Println("hero =", this)
}

func (this Hero) getName() string {
	return this.name
}

func (this *Hero) setName(val string)  {
	this.name = val
}
// 定义类结束

func main() {
	var hero Hero
	hero.name = "dao77777"
	hero.ad = 5
	hero.level = 12
	hero.show()
	fmt.Println(hero.getName())
	hero.setName("haha")
	fmt.Println(hero.getName())
}