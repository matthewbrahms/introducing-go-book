package main

import "fmt"

func main() {
	var celsius float64
	var fahrenheit float64
	fmt.Println("Enter a temperature in Fahrenheit: ")
	fmt.Scanf("%f", &fahrenheit)
	celsius = ((fahrenheit - 32) * 5/9)
	fmt.Println(fahrenheit, " degrees Fahrenheit is ", celsius, " degrees Celsius.")
}