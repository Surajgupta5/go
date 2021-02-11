package main

import "fmt"

func main() {
	y := 10
	a := 1
	for x := 0; x < y; x++ {
		fmt.Println("value of x is --> ", x)
	}

	fmt.Println("\nonly condition in for loop")
	for a < y {
		fmt.Println("value of x is --> ", a)
		a++
	}

	fmt.Println("\nusing break statement")

	for b := 1; b < y; b++ {

		fmt.Println("value of b is --> ", b)
		if b == 3 {
			break
		}

	}

	fmt.Println("\nusing continue statement")

	for b := 1; b < y; b++ {

		if b == 3 {
			continue
		}
		fmt.Println("value of b is --> ", b)

	}
}
