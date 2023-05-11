package main

import (
	"fmt"
	"time"
	"sync"
)

func worker(id int){
	fmt.Println("worker ", id, " starting")
	time.Sleep(time.Second)
	fmt.Println("worker ", id, " done")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}