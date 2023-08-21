package main

import (
	"errors"
	"fmt"
)

type Tom struct {
}
type Jim struct {
}

func main() {
	//空接口
	//var c interface{}
	var a Human
	var b Human
	t := Tom{}
	a = t
	fo(a)
	j := Jim{}
	b = j
	fo(b)
	//foo(1, 2)
	//fmt.Println(foo(1, 2))
	//fmt.Println(add(10, 0))
}

// 接口:接口是一组行为规范的集合
type Human interface {
	Say(int, int) int
}

func fo(h Human) {
	c := h.Say(3, 4)
	fmt.Println("c=", c)
}
func (Tom) Say(a int, b int) int {
	return a + b
}
func (Jim) Say(a int, b int) int {
	return a - b
}
func foo(a int, b int) int {
	//defer 要在方法内全部代码执行完毕后才可执行，后注册的defer先执行
	//defer后面是语句时，在注册时候就计算出来 如果是匿名函数，那么在函数计算时才执行
	d := a + b
	defer fmt.Println("111", d)
	fmt.Println(d)
	defer fmt.Println("222", d)
	defer func() {
		fmt.Println("333", d)
	}()
	d = 100
	return d
}
func add(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("number error")
	}
	//c := a + b
	d := a / b
	return d, nil
}
