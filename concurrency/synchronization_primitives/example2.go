package synchronization_primitives

//package main

import (
	"fmt"
	"sync"
)

func main() {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	v := 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()

			v++

			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(v)
}
