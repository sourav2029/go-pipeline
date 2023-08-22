package main

import (
	"fmt"
	"sync"
)

func main() {
	done := make(chan struct{})
	defer close(done)
	in := gen(done, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	c1 := sq(done, in)
	c2 := sq(done, in)
	out := merge(done, c1, c2)
	fmt.Println(<-out)
	fmt.Println(<-out)

	for n := range merge(done, c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}

}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}
