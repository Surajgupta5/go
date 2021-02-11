package main

import "fmt"

func main() {

	var x int
	x = 20
	if x > 50 {
		fmt.Println("False")
	} else if x > 10 {
		fmt.Println("yes x is greater than 10")
	} else {
		fmt.Println("Nothing")
	}
}
