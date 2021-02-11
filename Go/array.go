package main

import "fmt"

func main() {
	var n [10]int
	var i, j int
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("element[%d] = %d\n", j, n[j])
	}

	fmt.Println("for multidimmentional array")

	var a = [5][2]int{{1, 2}, {5, 4}, {7, 3}, {2, 3}, {9, 3}}
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
	// short method
	arr := []int{1, 2, 3, 5, 7, 23, 56}
	fmt.Println(arr)


	// slices

	slc := arr[]

}
