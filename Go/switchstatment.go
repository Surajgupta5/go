package main

import "fmt"

func main() {
	_result := "C"
	switch {
	case _result == "A":
		fmt.Println("very good grade")
	case _result == "B":
		fmt.Println("good grade")
	case _result == "C":
		fmt.Println("not good you can do it better")
	case _result == "D":
		fmt.Println("bad grade")
	default:
		fmt.Println("no case found")
	}
}
