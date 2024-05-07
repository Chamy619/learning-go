package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println("go1: do something")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("go2: do something")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("go3: do something")
	}()
	wg.Wait()

}
