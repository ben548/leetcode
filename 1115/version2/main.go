package main

import (
	"fmt"
)

var (
	ch1 = make(chan int, 1)
	ch2 = make(chan int, 1)
	stopChan = make(chan int)
	n = 3
)

func main()  {
	go first()
	go second()
	ch1<-1
	<- stopChan
}
func first()  {
	for i := 0; i < n; i++ {
		<- ch1
		fmt.Println("first")
		ch2 <- 1

	}
}

func second()  {
	for i := 0; i < n; i++ {
		<- ch2
		fmt.Println("second")
		ch1 <- 1
	}
	stopChan <- 1
}


