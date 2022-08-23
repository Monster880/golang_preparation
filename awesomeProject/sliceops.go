package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
}

func main() {
	var s []int
	//fmt.Println(s == nil)
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	//fmt.Println(s)
	s1 := []int{2, 4, 6, 8, 10}
	s1 = append(s1, 12)
	printSlice(s1)
	//printSlice(s2)

	s3 := make([]int, 16)
	s4 := make([]int, 10, 32)
	printSlice(s3)
	printSlice(s4)
	fmt.Println("Copying slice")
	copy(s3, s4)
	printSlice(s3)

	fmt.Println("Deleting elements from slice")
	//s3[:3] + s4[4:]
	s3 = append(s3[:3], s4[4:]...)
	printSlice(s3)

	fmt.Println("Popping from front")
	//front := s3[0]
	s3 = s3[1:]
	fmt.Println()
}
