package main

import "fmt"

func main() {
	mapFun := map[string]int{} //here string is the type for "key", it can be of any type maybe float or int.
	// here int is the type of value, maybe of difn datatype like key
	mapFun["A"] = 2
	mapFun["B"] = 234
	mapFun["C"] = 43
	mapFun["D"] = 5
	mapFun["E"] = 7
	mapFun["F"] = 1
	mapFun["G"] = 32

	//if we want to delete a value
	delete(mapFun, "C") // delete is the builtin function
	// mapFun is the name of variable
	// "C" is the key to delete particular value
	fmt.Println(mapFun)
}

// OutPut
//map[A:2 B:234 C:43 D:5 E:7 F:1 G:32]

// after using delete() function

// output will be
// map[A:2 B:234 D:5 E:7 F:1 G:32]
