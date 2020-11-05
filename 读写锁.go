package main

import (
	"fmt"
	"strconv"
	"sync"
)

type book struct {
	content int
	mu      *sync.Mutex
}

func newBook() *book {
	return &book{0, new(sync.Mutex)}
}

func write() {
	for {
		b.mu.Lock()
		b.content += 1
		b.mu.Unlock()
	}
}
func read() {
	a := 0
	for {
		fmt.Println(strconv.Itoa(a) + "-" + strconv.Itoa(b.content))
		// fmt.Println(strconv.Itoa(a) + b.content)
		a++
	}
}

var b = newBook()

func main() {
	go write()
	go read()
	select {}
}

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var mut sync.Mutex

// func printer(str string) {
// 	// mut.Lock()
// 	// defer mut.Unlock()
// 	for _, data := range str {
// 		fmt.Printf("%c", data)
// 		time.Sleep(800 * time.Microsecond)
// 	}
// 	fmt.Println()
// }
// func person1() {
// 	printer("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
// }
// func person2() {
// 	printer("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
// }
// func main() {
// 	go person1()
// 	go person2()
// 	for {
// 	}
// }

// //输出结果
// //world
// //hello
