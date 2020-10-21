package main

import "fmt"

var myPointer *int

func main() {
	fmt.Println(*myPointer)
}
func init() {
	myPointer, err := GenIntPointer()
	if err != nil {
		fmt.Println(err)
	}
	print(myPointer)
}
func print(myPointer *int) {
	fmt.Println(*myPointer)
}
func GenIntPointer() (*int, error) {
	i := 0
	return &i, nil
}

// panic: runtime error: invalid memory address or nil pointer dereference
// 问题在第11行，虽然在外面定义了一个全局变量myPointer，但是因为使用了 := 导致依然被当做局部变量处理。
// 于是init函数结束后局部变量指向的内存便被释放了，实际上并没有初始化myPointer
