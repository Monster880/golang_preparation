package main

//func main() {
//	s := "Yes我爱慕课网!"
//	//fmt.Printf("%X\n", s)
//	fmt.Println("Rune count: ", utf8.RuneCountInString(s))
//	for _, b := range s {
//		fmt.Printf("%c ", b)
//	}
//	fmt.Println()
//	for i, ch := range s { // ch is a rune
//		fmt.Printf("（%d %X）", i, ch)
//	}
//	fmt.Println()
//	fmt.Println("Rune count: ", utf8.RuneCountInString(s))
//	bytes := []byte(s)
//	for len(bytes) > 0 {
//		ch, size := utf8.DecodeRune(bytes)
//		bytes = bytes[size:]
//		fmt.Printf("%c ", ch)
//	}
//	fmt.Println()
//
//	for i, ch := range []rune(s) {
//		fmt.Printf("(%d %c)", i, ch)
//	}
//	fmt.Println()
//}
