package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("ShowA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("ShowB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher ShowB")
}
func main() {
	t := Teacher{}
	t.ShowA()
}

// 输出：
// ShowA
// ShowB
// 可以理解为go在调用内部方法时，会把逻辑预先copy进去。
