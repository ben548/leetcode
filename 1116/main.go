package main

import (
	"fmt"
	"sync"
)

var (
	ch1 = make(chan int)
	ch2 = make(chan int)
	ch3 = make(chan int)
	stopChan = make(chan int)
	wg sync.WaitGroup
	n = 5
)

func main()  {
	wg.Add(3)

	go zero()
	go odd()
	go even()

	for i := 1; i <= n; i++ {
		ch3 <- i
		<- stopChan
		if i % 2 == 0 {
			ch2 <- i
			<- stopChan
		} else {
			ch1 <- i
			<- stopChan
		}
	}

	close(ch1)
	close(ch2)
	close(ch3)

	wg.Wait()
}
func zero()  {
	defer wg.Done()
	for {
		_, ok := <- ch3
		if !ok {
			return
		}
		fmt.Print(0)
		stopChan <- 1
	}
}

func even()  {
	defer wg.Done()
	for {
		num, ok := <- ch2
		if !ok {
			return
		}
		fmt.Print(num)
		stopChan <- 1
	}
}

func odd()  {
	defer wg.Done()
	for {
		num, ok := <- ch1
		if !ok {
			return
		}
		fmt.Print(num)
		stopChan <- 1
	}
}


