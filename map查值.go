package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 0
	if v, ok := m["c"]; ok {
		fmt.Println("ok", v)
	} else {
		fmt.Println("not ok")
	}

}
