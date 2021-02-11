package main

import "fmt"

func main() {
	a := 20
	var ptr *int

	ptr = &a
	fmt.Println("address of a ", &a)
	fmt.Println("address stored in ptr ", ptr)
	fmt.Println("value of ptr ", *ptr)

	fmt.Println("\ndouble pointer\n")

	var dptr **int

	dptr = &ptr
	fmt.Println("value of a ", a)
	fmt.Println("value of ptr ", *ptr)
	fmt.Println("value of dptr ", **dptr)

	fmt.Println("\ntripple pointer\n")

	var tptr ***int

	tptr = &dptr
	fmt.Println("value of a ", a)
	fmt.Println("value of ptr ", *ptr)
	fmt.Println("value of dptr ", **dptr)
	fmt.Println("value of tptr ", ***tptr)

}

//Output

// address of a  0xc0000120b0
// address stored in ptr  0xc0000120b0
// value of ptr  20

// double pointer

// value of a  20
// value of ptr  20
// value of dptr  20

// tripple pointer

// value of a  20
// value of ptr  20
// value of dptr  20
// value of tptr  20
