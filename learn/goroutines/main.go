package main

/*

Take an integer N as input
Launch N goroutines, each adding their number to a slice
Return the slice of numbers once all goroutines complete
Handle concurrent access to the slice safely

Some hints if you need them:

You'll need a way to wait for all goroutines to complete
You'll need to protect access to the shared slice
The function signature should be: func ParallelCount(n int) []int
*/

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(ParallelCount(7))
}

func ParallelCount(n int) (res []int) {
	ch := make(chan int, n)

	var wg sync.WaitGroup

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()

			ch <- num
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	res = make([]int, 0, n)
	for num := range ch {
		res = append(res, num)
	}
	return res
}
