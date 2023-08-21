package main

import "fmt"

//空接口
//var c interface{}

// Human 接口:接口是一组行为规范的集合
type Human interface {
	Say(int, int) int
}
type Tom struct {
}
type Jim struct {
}

func (Tom) Say(a int, b int) int {
	return a + b
}
func (Jim) Say(a int, b int) int {
	return a - b
}
func test(h Human) {
	c := h.Say(3, 4)
	fmt.Println("c=", c)
}

func main() {
	var a Human
	var b Human
	t := Tom{}
	a = t
	test(a)
	j := Jim{}
	b = j
	test(b)
}
