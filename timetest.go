package main

import (
	"fmt"
	"time"
)

func main() {
	// pingTimeDefault := time.Now().Format("2006-01-02 15:04:05")
	// pingTimeCST := time.Now().Format("2006-01-02 15:04:05 CST")
	// pingTime := pingTimeCST[:len(pingTimeCST)-3]
	// fmt.Println(pingTime)
	// fmt.Println(pingTimeDefault)
	// fmt.Println(pingTimeCST)

	myTimeLocal := time.Now().UTC().Format("2006-01-02 15:04:05")
	local, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
	}
	myTime, _ := time.ParseInLocation("2006-01-02 15:04:05", myTimeLocal, local)
	fmt.Println("myTime", myTime.Format("2006-01-02 15:04:05"))
	// fmt.Println(myTime)
	// mytime2, _ := time.Parse("2006-01-02 15:04:05", time.Now().String())
	// fmt.Println("mytime2", mytime2)
	localCN, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
	}
	pingTime := time.Now().In(localCN).Format("2006-01-02 15:04:05")
	fmt.Println("pingTime", pingTime)
}
