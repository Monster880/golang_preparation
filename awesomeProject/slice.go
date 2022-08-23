package main

func updateSlice(s []int) {
	s[0] = 100
}

//func main() {
//	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
//	//s := arr[2:6]
//	//s1 := arr[:]
//	//updateSlice(s1)
//	//fmt.Println(s1)
//	//
//	//updateSlice(s)
//	//fmt.Println(s)
//	s1 := arr[2:6]
//	s2 := s1[3:5]
//	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
//	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
//	//fmt.Println(s1[3:6])
//	s3 := append(s2, 10)
//	s4 := append(s3, 11)
//	s5 := append(s4, 12)
//	fmt.Println("s3, s4, s5=", s3, s4, s5)
//	fmt.Println("arr=", arr)
//}
