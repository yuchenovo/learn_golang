package main

import (
	"fmt"
	"sync"
)

var (
	ch = make(chan int, 0)
)

// 该方法中 ch<-1 会直接死锁 而不是阻塞
func testDeadLock1() {
	ch <- 1
	fmt.Println("over")
}

// 该方法子协程一直在阻塞，main协程在休眠 但是并没有直接死锁 因为main协程对休眠后面的代码抱有期待 （<-ch） 如果用 ch <- 1 代替休眠 会直接死锁
func testDeadLock2() {
	go func() {
		ch <- 1
		fmt.Println("over")
	}()
	//time.Sleep(2 * time.Second)
	//<-ch
	ch <- 1
}

// 该方法main协程 在等待子协程Done 但是子协程阻塞无法Done 所以直接死锁
func testDeadLock3() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 1
	}()
	wg.Wait()
}

func testDeadLock4() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 1
	}()
	go func() {
		wg.Done()
	}()
	wg.Wait()
}

func main() {
	//testDeadLock2()
	//testDeadLock3()
	testDeadLock4()
}
