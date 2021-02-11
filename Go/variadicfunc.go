package main

import "fmt"

func main() {
	variadic(2, 4, 5, 6, 7)
}

func variadic(a ...int) { // from this we can pass multiple parameters in the func
	//variadic is the name of func we can right anything.

	for b := 0; b < len(a); b++ {

		fmt.Println(a[b])

	}
	fmt.Printf("%T\n%d", a, a) //%T is used for the type, whether it is int or float or anything

}
