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
	// struct是值类型。创建一样的struct,值相等
	// fmt.Println(A > B) //func can only be compared to nil. operator    > not defined on func
	fmt.Println("a" > "b")
	// a := "a"
	// fmt.Println(s1 == nil)

}

// 可排序的数据类型有3种，int、float、string
// 值类型：除了上述3种外还有：bool、array和struct，都是基数据类型，无法与nil比较
// 可比较的数据类型除了上述6种外，还有Pointer，Channel，interface这3个引用类型。而且引用类型可与nil比较
// 剩下3个引用类型：slice、map、function 是不可比较的数据类型包括，仅能和nil 做等值==比较

// 另外struct可不可比较，由成员存不存在不可比较类型决定

func A() {}
func B() {}
