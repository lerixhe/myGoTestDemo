package main

import "fmt"

func main() {
	var s1 = []int{2} // 初始化一个切片
	addOne(s1)        // 调用函数添加一个切片
	fmt.Println(s1)   // 输出一个值 [4]
	addOneFix(&s1)
	fmt.Println(s1) // 输出2个值 [8 1]

}

//定义一个函数，给切片添加一个元素
func addOne(s []int) {
	s[0] = 4         // 可以改变原切片值
	s = append(s, 1) // 扩容后分配了新的地址，原切片将不再受影响
	s[0] = 8
}

//定义一个函数，给切片添加一个元素
func addOneFix(s *[]int) {
	(*s)[0] = 4
	(*s) = append((*s), 1)
	(*s)[0] = 8
}
