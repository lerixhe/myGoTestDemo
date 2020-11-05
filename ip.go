package main

import "fmt"

func main() {

	now, per, rate := 1.0, 1.0, 1.0
	if now*rate > per || now/rate < per {
		fmt.Println("差距大")
	}
}
