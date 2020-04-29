package main

import "fmt"

func main() {
	var m string
	var n float32
	str:="aaaaaa 1.001 cccccc"
	fmt.Sscanf(str, "%s%v%s", &m, &n)
	fmt.Println(m,n)
}