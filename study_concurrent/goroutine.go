package main

import (
	"fmt"
	"time"
)

func fish() {
	//当子协程出现panic时，不影响main
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("panic happend: %s\n", err)
		}
	}()
	a, b := 3, 0
	fmt.Println(a, b)
	_ = a / b
	fmt.Println("fish01")
}

// main协程
func main() {
	//wg := sync.WaitGroup{}
	//wg.Add(2)
	//创建一个子协程来运行函数
	//go fish()
	//匿名函数创建
	//go func() {
	//	//对wg的值-1
	//	defer wg.Done()
	//	fmt.Println("fish")
	//}()
	//go func() {
	//	defer wg.Done()
	//	fmt.Println("fish1")
	//}()
	//Wait() 等待wg的值为0
	//wg.Wait()

	go fish()
	time.Sleep(1 * time.Second)
}
