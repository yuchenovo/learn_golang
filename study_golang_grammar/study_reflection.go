package main

import (
	"fmt"
	"reflect"
)

type Enum int

const (
	Zero Enum = 0
)

type cat struct {
}

func main() {
	var a int
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
	typeZ := reflect.TypeOf(Zero)
	fmt.Println(typeZ.Name(), typeZ.Kind())
	typeC := reflect.TypeOf(cat{})
	fmt.Println(typeC.Name(), typeC.Kind())

	//TODO
}
