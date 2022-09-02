package main

//func main() {
//	s := "abc[{2222]}ddd"
//	fmt.Println(check_brackets(s))
//}

func check_brackets(s string) bool {
	q := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == 40 || s[i] == 91 || s[i] == 123 {
			q = append(q, s[i])
		}
		if s[i] == 41 {
			if len(q) <= 0 || q[len(q)-1] != 40 {
				return false
			}
			q = q[:len(q)-1]
		}
		if s[i] == 93 {
			if len(q) <= 0 || q[len(q)-1] != 91 {
				return false
			}
			q = q[:len(q)-1]
		}
		if s[i] == 125 {
			if len(q) <= 0 || q[len(q)-1] != 123 {
				return false
			}
			q = q[:len(q)-1]
		}
	}
	return len(q) == 0
}

// 40 (
// 41 )
// 91 [
// 93 ]
// 123 {
// 125 }
