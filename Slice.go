package main

import "fmt"

func main() {

	Slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(Slice1)
	fmt.Println(len(Slice1))
	fmt.Println(cap(Slice1))

	var Slice2 []int
	Slice2 = append(Slice1, 6, 7)
	fmt.Println(len(Slice2))
	fmt.Println(cap(Slice2))
	Slice3 := make([]int, 32, 33)
	fmt.Println(len(Slice3))
	fmt.Println(cap(Slice3))
	fmt.Println(Slice3)

}
