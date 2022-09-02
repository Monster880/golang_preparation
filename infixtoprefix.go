package main

import "fmt"

//func main() {
//	priority()
//}

func priority() {
	s := "a你好cd"
	s = string([]rune(s)[:3])
	fmt.Println(s)
}
