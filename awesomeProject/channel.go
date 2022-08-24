package main

import (
	"time"
)

//func worker(id int, c chan int) {
//	for {
//		n, ok := <-c
//		fmt.Printf(" worker %d received %c\n", id, <-c)
//	}
//}

func chanDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		//go worker(i, channels[i])
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
	//n := <-c
	//fmt.Println(n)
}

//func bufferedChannel() {
//	c := make(chan int, 3)
//	go worker(0, c)
//	c <- 'a'
//	c <- 'b'
//	c <- 'c'
//	c <- 'd'
//	close(c)
//	time.Sleep(time.Millisecond)
//}
//
////func channelClose() {
////	c := make(chan int)
////}
//
//func main() {
//	chanDemo()
//	bufferedChannel()
//}
