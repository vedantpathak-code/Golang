package main

import "fmt"

func Max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func main() {
	var num1 int = 23
	var num2 int = 32

	fmt.Println(Max(num1, num2))
}