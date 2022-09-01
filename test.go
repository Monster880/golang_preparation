package main

import (
	"strconv"
)

//输入一个字符串，返回字符串表示的数学表达式的值。
//字符串里面只会包括正整数，+，- 两种运算
//比如 "20+10-5"，返回25
//备注：没有括号，不考虑溢出

//func main() {
//	s := "20+10-5"
//	res := computeMathString(s)
//	fmt.Println(res)
//}

func computeMathString(s string) int64 {
	str := make([]string, 0)
	operator := make([]byte, 0)
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 43 || s[i] == 45 {
			str = append(str, s[start:i])
			operator = append(operator, s[i])
			start = i + 1
		} else if i == len(s)-1 {
			str = append(str, s[i:])
		}
	}
	num := int64(0)
	number := 0
	for i := 0; i < len(operator); i++ {
		if operator[i] == 43 {
			if number == 0 {
				num1, _ := strconv.ParseInt(str[number], 10, 64)
				number++
				num2, _ := strconv.ParseInt(str[number], 10, 64)
				number++
				num = num + num1 + num2
			} else {
				num1, _ := strconv.ParseInt(str[number], 10, 64)
				number++
				num = num + num1
			}
		}
		if operator[i] == 45 {
			num1, _ := strconv.ParseInt(str[number], 10, 64)
			number++
			num = num - num1
		}
	}

	//fmt.Println(str)
	//fmt.Println(operator)
	//fmt.Println(num)
	return num
}
