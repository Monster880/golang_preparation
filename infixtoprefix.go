package main

import "fmt"

func main() {
	fmt.Println("(3+4)*5-6")
	fmt.Println(in2pre([]rune("(3+4)*5-6")))
}

func priority(s string) int {
	if s == "+" || s == "-" {
		return 1
	} else if s == "*" || s == "/" {
		return 2
	} else if s == ")" {
		return 3
	} else if s == "(" {
		return 0
	}
	return -1
}

func in2pre(s []rune) string {
	res := ""
	num := make([]rune, 0)
	operator := make([]rune, 0)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			num = append(num, s[i])
		} else if len(operator) == 0 || priority(string(s[i])) >= priority(string(operator[len(operator)-1])) {
			operator = append(operator, s[i])
		} else {
			for len(operator) > 0 && string(operator[len(operator)-1]) != ")" {
				num = append(num, operator[len(operator)-1])
				operator = operator[:len(operator)-1]
			}
			if string(s[i]) == "(" {
				operator = operator[:len(operator)-1]
			} else {
				operator = append(operator, s[i])
			}
		}
	}
	for len(operator) > 0 {
		num = append(num, operator[len(operator)-1])
		operator = operator[:len(operator)-1]
	}
	for len(num) > 0 {
		res += string(num[len(num)-1])
		num = num[:len(num)-1]
	}
	return res
}

//s := "a你好cd"
//s = string([]rune(s)[:3])
//fmt.Println(s)
