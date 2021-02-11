package main

import "fmt"

type Students struct {
	name      string
	id        int
	contactno int
}

func main() {
	var s1, s2 Students

	s1.contactno = 9876543312
	s1.id = 101
	s1.name = "samar"

	s2.contactno = 9756742752
	s2.id = 102
	s2.name = "karan"

	fmt.Println("name of Student1 is", s1.name)
	fmt.Println("id of Student1 is", s1.id)
	fmt.Println("contactno. of Student1 is", s1.contactno)
	fmt.Println("name of Student1 is", s2.name)
	fmt.Println("id of Student1 is", s2.id)
	fmt.Println("contactno.of Student1 is", s2.contactno)

	fmt.Println("\nstructure as a function arguments\n")

	printStudent(s1)
	printStudent(s2)

}

func printStudent(std Students) {
	fmt.Println("name of Student is", std.name, "id of Student is", std.id, "contactno. of Student is", std.contactno, "\n")

}
