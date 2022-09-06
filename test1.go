package main

//Input: candidates = [10,1,2,7,6,1,5], target = 8
//Output:
//[[1,1,6], [1,2,5], [1,7], [2,6]]

var res [][]int
var target int

//func main() {
//	candidates := []int{10, 1, 2, 7, 6, 1, 5}
//	target = 8
//	sort.Ints(candidates)
//	res = make([][]int, 0)
//	dfs1(candidates, []int{}, 0, 0)
//	fmt.Println(res)
//}

func dfs1(candidates []int, temp []int, sum int, index int) {
	if len(temp) > 0 && sum == target {
		arr := make([]int, len(temp))
		copy(arr, temp)
		res = append(res, arr)
		return
	}
	if target < 0 {
		return
	}
	for i := index; i < len(candidates); i++ {
		if i-1 >= index && candidates[i-1] == candidates[i] {
			continue
		}
		temp = append(temp, candidates[i])
		dfs1(candidates, temp, sum+candidates[i], i+1)
		temp = temp[:len(temp)-1]
	}
}
