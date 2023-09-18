package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
 * 存在并发问题，无法正确按顺序输出
 */
func sync01() {
	var a = 0
	for i := 0; i < 1000; i++ {
		go func(idx int) {
			a += 1
			fmt.Println(a)
		}(i)
	}
	time.Sleep(time.Second)
}

/*
添加互斥锁
*/
func sync02() {
	var a = 0
	var lock sync.Mutex
	for i := 0; i < 1000; i++ {
		go func(idx int) {
			lock.Lock()
			defer lock.Unlock()
			a += 1
			fmt.Printf("goroutine %d, a=%d\n", idx, a)
		}(i)
	}
	time.Sleep(time.Second)
}

/*
同时只能有一个 goroutine 能够获得写锁定；
同时可以有任意多个 gorouinte 获得读锁定；
同时只能存在写锁定或读锁定（读和写互斥）。
*/
var count int
var rw sync.RWMutex

func read(n int, ch chan struct{}) {
	rw.RLock()
	fmt.Printf("goroutine %d 进入读操作...\n", n)
	v := count
	fmt.Printf("goroutine %d 读取结束，值为：%d\n", n, v)
	rw.RUnlock()
	ch <- struct{}{}
}
func write(n int, ch chan struct{}) {
	rw.Lock()
	fmt.Printf("goroutine %d 进入写操作...\n", n)
	v := rand.Intn(1000)
	count = v
	fmt.Printf("goroutine %d 写入结束，新值为：%d\n", n, v)
	rw.Unlock()
	ch <- struct{}{}
}
func sync03() {
	ch := make(chan struct{}, 10)
	for i := 0; i < 5; i++ {
		go read(i, ch)
	}
	for i := 0; i < 5; i++ {
		go write(i, ch)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
}
func main() {
	//sync01()
	//sync02()
	sync03()
}
