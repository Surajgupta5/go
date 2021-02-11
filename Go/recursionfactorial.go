package main

import "fmt"

func main() {
	i := 20
	fmt.Println("Factorial of", i, "is", factorial(i))

}

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}
