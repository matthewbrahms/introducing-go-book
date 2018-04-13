package main

import "fmt"

func main() {
	fmt.Println(fibonacci(7))
}

func fibonacci(n int) int  {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibonacci(n-1) + fibonacci(n-2)
	}
}