package main

import "fmt"

type S struct {
	i int
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := ch1
	fmt.Println(ch1 == ch2)
	fmt.Println(ch1 == ch3)
	// false
	// true
	// channel 是引用类型，但语法上可比较。用make创建意味着底层指针不同，创建一样的channel不相等
	s1 := S{1}
	s2 := S{1}
	s3 := s1
	fmt.Println(s1 == s2)
	fmt.Println(s1 == s3)
	// true
	// true
	// struct是类型。创建一样的struct,值相等
}

// 可排序的数据类型有三种，Integer，Floating-point，和String
// 可比较的数据类型除了上述三种外，还有Boolean，Complex，Pointer，Channel，Interface和Array
// 不可比较的数据类型包括，Slice, Map, 和Function
// 另外struct可不可比较，由成员存不存在不可比较类型决定
