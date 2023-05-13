package main

import (
	"fmt"
	"sync"
)

var (
	ch1 = make(chan int, 1)
	ch2 = make(chan int, 1)
	wg sync.WaitGroup
	n = 3
)

func main()  {
	wg.Add(2)
	go first()
	go second()
	ch1<-1
	wg.Wait()
}
func first()  {
	defer wg.Done()
	//wg.Add(1)
	for i := 0; i < n; i++ {
		<- ch1
		fmt.Println("first")
		ch2 <- 1

	}
	//wg.Done()
}

func second()  {
	defer wg.Done()
	//wg.Add(1)
	for i := 0; i < n; i++ {
		<- ch2
		fmt.Println("second")
		ch1 <- 1
	}
	//wg.Done()
}


