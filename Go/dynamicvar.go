//var x float64 = 20.12  same for int64 too
// x := 50          |
// x := 12.30       |---->dynamic variable declaration
// x := "asdfgdas"  |

package main

import "fmt"

func main() {
	var a float64 = 12.40
	x := 50
	fmt.Println("the value of a and x :- ", a, x)
	fmt.Printf("x is of type %T \n", x)

	//for mixed type variable
	var b, c, d = 2, 3.45, "Hello World"
	fmt.Println("the value of b and c :- ", b, c)
	fmt.Println("the value of d :- ", d)
	fmt.Printf("d is of type %T", d)

}
