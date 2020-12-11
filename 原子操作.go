package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// 多个gorountine的计数器
	var counter uint64 = 5000000
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100000; j++ {
				atomic.AddUint64(&counter, ^uint64(0))
				// counter++
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("counter:", counter)
}

// 有符号数中负数求补码：反码+1得到补码
// 补码的特性是补码+原码=模，而模就是溢出值。所以原码+反码=溢出0,而模=临界值+1=溢出0
// 所以原码+反码+1=溢出0=模

// 无符号数没有补码，反码，原码之分，一个无符号数 通过补码的逆运算，可以得到一个负数（相反数）的同等作用的实值
// 但可以推出类似的概念。一个无符号数可以认为是，另一个符号位为1的无符号数的正向补码，求这个符号位为1的无符号数操作。
