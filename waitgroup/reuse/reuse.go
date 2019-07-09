package main

import (
	"fmt"
	"sync"

	"github.com/andrii-stasiuk/go-testing/waitgroup/calc"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	var x, y int = 10, 5
	go func() {
		fmt.Println(calc.Add(x, y))
		wg.Done()
	}()
	go func() {
		fmt.Println(calc.Subtruct(x, y))
		wg.Done()
	}()
	wg.Wait()
}
