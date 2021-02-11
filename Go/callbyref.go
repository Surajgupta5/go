package main

import "fmt"

func main() {
	a, b := 10, 20
	fmt.Println("before swap value of a is ", a)
	fmt.Println("before swap value of b is ", b)
	swap(&a, &b)
	fmt.Println("after swap value of a is ", a)
	fmt.Println("after swap value of b is ", b)

}
func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
