package main

import "fmt"

func main0() {
	c := make(chan string, 10)
	close(c)
	c <- "test"
}

// 向一个关闭的chan发送数据，panic

func main1() {
	c := make(chan string, 10)
	close(c)
	for a := range c {
		fmt.Printf("receive string %s\n", a)
	}
}

// range遍历一个已经关闭的chan，直接结束。

func main() {
	c := make(chan string, 10)
	close(c)
	for {
		select {
		case s := <-c:
			fmt.Printf("receive string %s\n", s)
		default:
			fmt.Println("no data")
		}
	}
}

// for select，可以不断从关闭的管道中读出零值。
