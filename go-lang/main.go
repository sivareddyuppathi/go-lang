package main

import (
	"fmt"
	"sort"
)

func main() {
	// Strings
	strs := []string{"c", "a", "b", "x", "d", "A"}
	fmt.Println(strs)                  // before sort
	sort.Strings(strs)                 // sorting
	fmt.Println(strs)                  //after sort
	a := sort.SearchStrings(strs, "B") //search for string and returns index else -1
	fmt.Println("Search: ", a)         //index printing

	//Integers
	ints := []int{1, 2, 4, 3, 1}
	fmt.Println(ints)              // before sort
	sort.Ints(ints)                //sorting
	fmt.Println(ints)              //after sort
	s1 := sort.IntsAreSorted(ints) // sorted or not if sorted returns true else false
	fmt.Println(s1)                //sorted or not
	s2 := sort.SearchInts(ints, 7) // if not avilable in array returns next immediate index..
	fmt.Println(s2)

	// fmt.Println(person{"Bob", 20})
	// fmt.Println(person{name: "Bob", id: 20})
	// fmt.Println(person{id: 3})
}

// type person struct {
// 	name string
// 	id   int
// }

// func xyz(x, y int) int {
// 	return x + y
// }
