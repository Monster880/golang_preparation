package main

import (
	"fmt"
	"time"
)

//golang题目:  异步日志收集器，实现日志数量达到100或者超过1s就上报

func main() {
	data := make(chan int)
	arr := make([]int, 0)
	canQuit := make(chan bool)

	go func() {
		flag := false
		start_time := time.Now()
		num := 0
		for d := range data {
			num++
			fmt.Println(num)
			arr = append(arr, d)
		}
		duration := time.Since(start_time)
		if duration == 60 || num == 100 {
			flag = true
		}
		canQuit <- flag
	}()

	for i := 0; i < 200; i++ {
		data <- i
	}
	if <-canQuit {
		close(data)
	}
}
