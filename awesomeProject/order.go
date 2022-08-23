package main

import (
	"fmt"
	"sort"
	"strconv"
	// "strings"
)

func main() {
	intSlice := []int{8, 90, 3, 303, 56, 74}
	sort.Slice(intSlice, func(i, j int) bool {
		a := strconv.Itoa(intSlice[i])
		b := strconv.Itoa(intSlice[j])
		return a+b > b+a
	})
	fmt.Println(intSlice)
	//res := strings.Join(intSlice, "")
	//fmt.Println(res)
	str := ""
	for i := 0; i < len(intSlice); i++ {
		str += strconv.Itoa(intSlice[i])
	}
	fmt.Println(str)
}
