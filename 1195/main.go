package main

import (
	"fmt"
	"sync"
)

var (
	ch1 = make(chan int, 1)
	ch2 = make(chan int, 1)
	ch3 = make(chan int, 1)
	ch4 = make(chan int, 1)
	stopChan = make(chan int, 1)
	wg sync.WaitGroup
	n = 15
)

func main()  {
	wg.Add(4)

	go number()
	go fizz()
	go buzz()
	go fizzbuzz()

	for i := 1; i <= n; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			ch4 <- 1
			<- stopChan
		} else if i % 3 == 0 {
			ch2 <- 1
			<- stopChan
		} else if i % 5 == 0 {
			ch3 <- 1
			<- stopChan
		} else {
			ch1 <- i
			<- stopChan
		}
	}

	close(ch1)
	close(ch2)
	close(ch3)
	close(ch4)

	wg.Wait()
}
func number()  {
	defer wg.Done()
	for {
		num, ok := <- ch1
		if !ok {
			return
		}
		fmt.Println(num)
		stopChan <- 1
	}
}

func fizz()  {
	defer wg.Done()
	for {
		_, ok := <- ch2
		if !ok {
			return
		}
		fmt.Println("fizz")
		stopChan <- 1
	}
}

func buzz()  {
	defer wg.Done()
	for {
		_, ok := <- ch3
		if !ok {
			return
		}
		fmt.Println("buzz")
		stopChan <- 1
	}
}

func fizzbuzz()  {
	defer wg.Done()
	for {
		_, ok := <- ch4
		if !ok {
			return
		}
		fmt.Println("fizzbuzz")
		stopChan <- 1
	}
}


