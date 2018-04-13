package main

import "fmt"

func main() {
	var feet float64
	var meters float64
	fmt.Println("Enter a number of feet to convert to meters:")
	fmt.Scanf("%f", &feet)
	meters = feet * 0.3048
	fmt.Println(feet, " feet is ", meters, " meters!")
}
