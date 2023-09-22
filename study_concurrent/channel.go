package main

import (
	"fmt"
	"sync"
)

/*
管道关键字chan
<- 指向管道代表插入元素 ch <- 1
<- 背离管道代表取出元素 a := <- ch
*/
func test() {
	//管道容量可为0
	ch := make(chan int, 8)
	ch <- 1
	a := <-ch
	fmt.Println(a)
}
func main() {
	ch := make(chan int, 100)
	wg := sync.WaitGroup{}
	wg.Add(2)
	//生产者1
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	//生产者2
	go func() {
		defer wg.Done()
		for i := 10; i < 20; i++ {
			ch <- i
		}
	}()
	//wg2 := sync.WaitGroup{}
	//wg2.Add(1)
	//消费者
	mc := make(chan struct{}, 0)
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				break
			} else {
				fmt.Println("consumer:", a)
			}
		}
		mc <- struct{}{}
		//wg2.Done()
	}()
	wg.Wait()
	close(ch)
	//wg2.Wait()
	<-mc
}
