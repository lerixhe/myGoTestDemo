package main

import (
	"fmt"
)

type MyStruct struct {
	Name string
	age  int
	sex  string
}

func main() {
	a := MyStruct{"xiaoming", 20, "m"}
	b := MyStruct{"xiaohong", 20, ""}
	c := MyStruct{"xiaoming", 20, ""}
	fmt.Println(a == b, a == c)
}
