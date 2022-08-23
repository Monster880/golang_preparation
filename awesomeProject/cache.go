package main

import "fmt"

//func main() {
//	//var arr1 [5]int
//	//arr2 := [3]int{1,3,5}
//	arr3 := [...]int{2, 4, 6, 8, 10}
//	//var grid [4][5]int
//
//	//printArray(arr1)
//	//printArray(arr3)
//	//arr3[0] = 10
//
//	printArray(&arr3)
//}

func printArray(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
