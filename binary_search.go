package main

//func main() {
//	arr := []int{1, 2, 3, 5, 6}
//	num := 6
//	index := binary_search(arr, num)
//	fmt.Println(index)
//}

func binary_search(arr []int, num int) int {
	left := 0
	right := len(arr)
	for left != right {
		mid := (left + right) >> 1
		if num < arr[mid] {
			right = mid
		} else if num >= arr[left] {
			left = mid + 1
		}
	}
	return left - 1
}
