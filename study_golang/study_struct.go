package main

import "fmt"

type Point struct {
	X int
	Y int
}
type People struct {
	name  string
	child *People
}
type Cat struct {
	Color string
	Name  string
}
type innerS struct {
	in1 int
	in2 int
}

// 内嵌结构体
type outerS struct {
	b      int
	c      float32
	int    // anonymous field
	innerS //anonymous field
}

func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

func NewCatByColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}

func main() {
	var p Point
	p.X = 10
	p.Y = 20
	//fmt.Println(p.X, p.Y)
	relation := &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}
	fmt.Println(relation.name, relation.child.name, relation.child.child.name)
	//fmt.Println(NewCatByColor("blue").Color)
	//fmt.Println(NewCatByName("Brain").Name)
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10
	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)
	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Printf("%+v\n", outer2)
}
