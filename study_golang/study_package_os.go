package main

import (
	"fmt"
	"os"
)

func open1() {
	file, err := os.Open("study_golang/1.txt")
	if err != nil {
		fmt.Println("open failed:", err)
		return
	}
	defer file.Close()
	content := make([]byte, 100)
	//read 为读取到的长度
	read, err := file.Read(content)
	if err != nil {
		fmt.Println("read failed:", err)
		return
	}
	fmt.Println(content[0:read])
	fmt.Println(content)
	fmt.Println(string(content[0:read]))
}

/*
0777表示：创建了一个普通文件，所有人拥有所有的读、写、执行权限
0666表示：创建了一个普通文件，所有人拥有对该文件的读、写权限，但是都不可执行
0644表示：创建了一个普通文件，文件所有者对该文件有读写权限，用户组和其他人只有读权限，都没有执行权限
*/
func open2() {
	file, err := os.OpenFile("study_golang/2.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open failed:", err)
		return
	}
	defer file.Close()
	content := "9月15日"
	n, err := file.Write([]byte(content))
	if err != nil {
		fmt.Println("write failed:", err)
	} else {
		fmt.Println(n)
	}
}
func main() {
	//open1()
	open2()
}
