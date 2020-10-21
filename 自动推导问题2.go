package main

import (
	"errors"
	"fmt"
)

var Err1 = errors.New("err1")
var Err2 = errors.New("err2")
var Err3 = errors.New("err3")

func main() {
	a, err := GenErr1()
	b, err := GenErr2()
	c, err := GenErr3()
	// d, err := GenNoErr()
	_, err = GenNoErr()
	if err != nil {
		fmt.Println(a, b, c, err)

	}
	fmt.Println("no error")

}

// no error
// 很显然，偷懒只检查最后一个err的操作是有问题的。err仅保留最后一次的赋值结果

func GenErr1() (int, error) {
	return 1, Err1
}
func GenErr2() (int, error) {
	return 2, Err2
}
func GenErr3() (int, error) {
	return 3, Err3
}
func GenNoErr() (int, error) {
	return 0, nil
}
