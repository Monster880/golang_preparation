package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args"+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b int) (int, int) {
	return b, a
}

//func main() {
//	fmt.Println(eval(3, 4, "*"))
//	if q, err := eval(3, 4, "/"); err != nil {
//		fmt.Println(q)
//	} else {
//		fmt.Println(err)
//	}
//	fmt.Println(apply(func(a, b int) int { return int(math.Pow(float64(a), float64(b))) }, 3, 4))
//	fmt.Println(sum(1, 2, 3, 4, 5))
//	a, b := 3, 4
//	a, b = swap(a, b)
//	fmt.Println(a, b)
//}

//
//func longestConsecutive(nums []int) {
//	sort.Ints(nums)
//}
