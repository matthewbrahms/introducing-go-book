package main

import "fmt"

func main() {
	fmt.Println(max(1,2,3,4,5,))
}

func max(x ...int) int  {
	var max int
	for i, v := range x {
		if i == 0 || v > max {
			max = v
		}
	}
	return max
}