package main

import (
	"fmt"
	"sync"
)

var (
	ch1 = make(chan int)
	ch2 = make(chan int)
	wg sync.WaitGroup
)

func main()  {
	water := "HOHHHOOHHOOHHHH"

	l := len(water)

	wg.Add(2)

	go H(l / 3 * 2)
	go O(l / 3)

	ch2 <- 1
	ch2 <- 1

	wg.Wait()
}
func H(n int)  {
	defer wg.Done()
	for i := 1; i <= n; i++ {
		<- ch2
		fmt.Print("H")
		if i % 2 == 0 {
			ch1 <- 1
		}
		if i == n {
			return
		}
	}
}

func O(n int)  {
	defer wg.Done()
	for i := 1; i <= n; i++ {
		<- ch1
		fmt.Print("O")
		if i == n {
			close(ch1)
			close(ch2)
			return
		}
		ch2 <- 1
		ch2 <- 1
	}
}


