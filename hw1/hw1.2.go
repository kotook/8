package main

import (
	"hw1/dir"
	"sync"
)

func main() {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	resultPath := dir.Dir()
	wg.Add(1000000)
	for i := 0; i < 1000000; i++ {

		go func(i int) {
			mu.Lock()
			dir.Create(resultPath, i)
			mu.Unlock()
			wg.Done()
		}(i)

	}
	wg.Wait()
}
