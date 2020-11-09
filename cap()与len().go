package main

import "fmt"

func main() {
	array := [5]int{1, 1, 1, 1, 1}
	slice := array[2:4]
	channel := make(chan int, 4)
	channel <- 1
	fmt.Println(len(array))
	fmt.Println(cap(array))
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(len(channel))
	fmt.Println(cap(channel))

}

// 数组、切片、管道都有len()和cap(),对于管道，len表示内部元素个数，cap表示缓冲区大小
