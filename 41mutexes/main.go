package main

import (
	"fmt"
	"sync"
)

type Containers struct {
	mu      sync.Mutex
	counter map[string]int
}

func (c *Containers) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter[name]++
}

func main(){
	counterMap := map[string]int{"a": 0, "b": 0}
	c := Containers{counter: counterMap}

	var wg sync.WaitGroup

	doInc := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}
	wg.Add(3)
	go doInc("a", 10000)
	go doInc("a", 10000)
	go doInc("b", 10000)
	wg.Wait()
	fmt.Println(c.counter)
}

