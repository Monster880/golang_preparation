package main

import "fmt"

var res1 []int

//func main() {
//	arr1 := []int{2, 3, 4}
//	arr2 := []int{0, 6, 8}
//	arr3 := []int{1, 5, 7}
//	res1 = make([]int, 0)
//	append3sortedArr(arr1, arr2, arr3)
//	fmt.Println(res1)
//}

func append3sortedArr(arr1 []int, arr2 []int, arr3 []int) {
	index1, index2, index3 := 0, 0, 0
	for index1 < len(arr1) || index2 < len(arr2) || index3 < len(arr3) {
		x, y, z := 100, 100, 100
		if index1 < len(arr1) {
			x = arr1[index1]
		}
		if index2 < len(arr2) {
			y = arr2[index2]
		}
		if index3 < len(arr3) {
			z = arr3[index3]
		}
		temp := min1(x, y, z)
		if temp == 1 {
			res1 = append(res1, arr1[index1])
			index1++
		}
		if temp == 2 {
			res1 = append(res1, arr2[index2])
			index2++
		}
		if temp == 3 {
			res1 = append(res1, arr3[index3])
			index3++
		}
		fmt.Println(res1)
	}
}

func min1(x, y, z int) int {
	if x < y && x < z {
		return 1
	} else if y < x && y < z {
		return 2
	}
	return 3
}
