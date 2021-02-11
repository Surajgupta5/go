package main

import "fmt"

func main() {
	a, b := 10, 20
	fmt.Println("before swap value of a is ", a)
	fmt.Println("before swap value of b is ", b)
	swapp(a, b)

}
func swapp(p int, q int) {
	var temp int
	temp = p
	p = q
	q = temp
	fmt.Println("after swap value of a is ", p)
	fmt.Println("after swap value of b is ", q)

}
