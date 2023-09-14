package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("1.txt")
	if err != nil {
		fmt.Println("failed:", err)
		return
	}
	content := make([]byte, 100, 100)
	read, err := file.Read(content)
	if err != nil {
		return
	}
	fmt.Println(read)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
}
