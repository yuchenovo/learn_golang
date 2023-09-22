package main

import (
	"fmt"
	"sort"
)

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

	names := []string{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}

	sort.Strings(names)

	//遍历打印结果
	for _, v := range names {
		fmt.Printf("%s\n", v)
	}
}
