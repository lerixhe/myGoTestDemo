package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	age  int
	sex  string
}
type People struct {
	Stus []*Student
}

func main() {
	a := Student{"xiaoming", 20, "m"}
	b := Student{"xiaoming", 20, "m"}
	c := Student{"xiaohong", 21, "m"}
	d := Student{"xiaohong", 21, "m"}
	stus0 := []*Student{&a, &c}
	stus1 := []*Student{&b, &d}
	p1 := People{stus0}
	p2 := People{stus1}
	p1p := &p1
	p2p := &p2
	fmt.Println(p1, p2)
	fmt.Println(p1p, p2p)
	// fmt.Println(p1 == p2) invalid operation: p1 == p2 (struct containing []Student cannot be compared)
	fmt.Println(reflect.DeepEqual(p1, p2))
	fmt.Println(reflect.DeepEqual(p1p, p2p))
}

// 结论：
// 1.切片只能和nil比较
// 2.deepequal不仅可以深度比较普通切片，也可以深度比较指针组成的切片,深度比较结构体指针。即无需考虑指针的存在，即可正常处理。
