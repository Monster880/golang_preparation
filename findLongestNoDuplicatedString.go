package main

//func main() {
//	var s = "abcabcbb"
//	fmt.Println(findLongestNoDuplicatedLength(s))
//}

func findLongestNoDuplicatedLength(s string) int {
	m := make(map[byte]bool, 0)
	end := 0
	res := 0
	for start := 0; start < len(s); start++ {
		for end < len(s) && !m[s[end]] {
			m[s[end]] = true
			end++
		}
		//fmt.Println(end)
		res = max(res, end-start)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
