package main

import (
	"fmt"
	"sync"
)

var (
	ch1 = make(chan int)
	ch2 = make(chan int)
	ch3 = make(chan int)
	wg sync.WaitGroup
)

func main()  {
	wg.Add(3)
	go first()
	go second()
	go third()
	ch1<-1
	wg.Wait()
}
func first()  {
	<- ch1
	fmt.Println("first")
	ch2 <- 1
	wg.Done()
}

func second()  {
	<- ch2
	fmt.Println("second")
	ch3 <- 1
	wg.Done()
}

func third()  {
	<- ch3
	fmt.Println("third")
	//ch1<- 1
	wg.Done()
}


