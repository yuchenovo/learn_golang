package main

import (
	"errors"
	"fmt"
)

func dev(b int) error {
	if b == 0 {
		return errors.New("number error")
	}
	return nil
}
func nine() {
	for a := 1; a < 10; a++ {
		for b := 1; b < 10; b++ {
			fmt.Print(a*b, "\t")
		}
		fmt.Println()
	}
}
func cha() {
	c := make(chan int)

	go func() {

		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
func Accumulate(value int) func() int {

	// 返回一个闭包
	return func() int {

		// 累加
		value++

		// 返回一个累加值
		return value
	}
}
func acc() {
	// 创建一个累加器, 初始值为1
	accumulator := Accumulate(10)
	// 累加1并打印
	fmt.Println(accumulator())
	fmt.Println(accumulator())

	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator)

	// 创建一个累加器, 初始值为1
	accumulator2 := Accumulate(10)

	// 累加1并打印
	fmt.Println(accumulator2())

	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator2)
}
func main() {
	//err != nil 才是 if 的判断表达式
	//if err := dev(0); err != nil {
	//}
	//nine()
	//cha()
	err := dev(0)
	if err != nil {
		goto exit
	}
exit:
	fmt.Println(err)
}
