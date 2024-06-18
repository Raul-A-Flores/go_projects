package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func dowork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func main() {

	// Channel is unbuffered and requires to be buffered or have a go routine
	// to send data
	//dataChan := make(chan int)

	// buffered channel, greater buffer more channels
	dataChan := make(chan int, 1)

	dataChan <- 8439
	//go func() {
	//	dataChan <- 8439
	//}()

	n := <-dataChan

	fmt.Println(n)

	dataChan2 := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := dowork()
				dataChan2 <- result
			}()
		}
		wg.Wait()
		close(dataChan2)
	}()

	for d := range dataChan2 {
		fmt.Println(d)
	}

}
