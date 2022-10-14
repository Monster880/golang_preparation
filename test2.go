package main

import (
	"fmt"
)

func main() {
	var graph = [6][6]int{{0, 1, 1, 0, 0, 0}, {1, 0, 2, 0, 1, 2}, {1, 2, 0, 1, 0, 0}, {0, 0, 1, 0, 0, 1}, {0, 1, 0, 0, 0, 1}, {0, 2, 0, 1, 1, 0}}
	res := make([][]int, 0)
	temp := []int{0}
	var dfs func(int)
	target := 10
	dfs = func(start int) {
		if start == 0 && target == 0 {
			res = append(res, append([]int(nil), temp...))
			return
		}
		for end := 0; end < 6; end++ {
			if graph[start][end] > 0 {
				temp = append(temp, end)
				graph[start][end]--
				graph[end][start]--
				target--
				dfs(end)
				temp = temp[:len(temp)-1]
				graph[start][end]++
				graph[end][start]++
				target++
			}
		}
	}
	dfs(0)
	for i := 0; i < len(res); i++ {
		temp := make([]rune, 0)
		for j := 0; j < len(res[i]); j++ {
			if res[i][j] == 0 {
				temp = append(temp, 'A')
			}
			if res[i][j] == 1 {
				temp = append(temp, 'B')
			}
			if res[i][j] == 2 {
				temp = append(temp, 'C')
			}
			if res[i][j] == 3 {
				temp = append(temp, 'D')
			}
			if res[i][j] == 4 {
				temp = append(temp, 'E')
			}
			if res[i][j] == 5 {
				temp = append(temp, 'F')
			}
		}
		outString := string(temp[0])
		for m := 1; m < len(res[i]); m++ {
			outString += string(temp[m])
		}
		fmt.Println(outString)
	}
	return
}
