package main

//func main() {
//	m := make(map[int]int, 0)
//	m[5] = 30
//	m[20] = 15
//	m[50] = 10
//	m[100] = 9
//	m[500] = 8
//	m[1000] = 7
//	m[2000] = 6
//	m[3000] = 5
//	m[4000] = 4
//	m[5000] = 3
//	m[6000] = 2
//	arr := make([]int, 0)
//	arr = append(arr, []int{5, 20, 50, 100, 500, 1000, 2000, 3000, 4000, 5000, 6000}...)
//	res := compute_cost(m, arr, 21)
//	fmt.Println(res)
//}

func compute_cost(m map[int]int, arr []int, num int) int {
	res := 0
	pre_k := 0
	for _, v := range arr {
		if num >= v {
			res += (v - pre_k) * m[v]
		} else {
			res += (num - pre_k) * m[v]
			return res
		}
		pre_k = v
	}
	if num >= 6000 {
		res += (num - 6000) * 2
	}
	return res
}
