package main

import "fmt"

func main() {
	a := 10
	b := 20
	var res int
	res = max(a, b)
	fmt.Println("max value is -->", res)
}

func max(num1 int, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1

	} else {
		result = num2
	}
	return result
}
