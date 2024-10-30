/*
Your task is to implement ParallelWordFilter that:

Takes a slice of strings and a minimum length
Processes each word in parallel using goroutines
Returns a slice containing only words >= minLength
Maintains the original order of the words
Uses proper concurrency patterns (WaitGroup, defer, etc.)
*/

package filter

import (
	"sync"
)

type result struct {
	word     string
	position int
	valid    bool
}

func ParallelWordFilter(stringArray []string, minLength int) []string {
	var wg sync.WaitGroup
	ch := make(chan result, len(stringArray)) // buffered array
	validResults := make(map[int]string)

	for i, s := range stringArray {
		wg.Add(1)
		go func(s string, pos int) {
			defer wg.Done()

			if len(s) >= minLength {
				ch <- result{
					word:     s,
					position: pos,
					valid:    true,
				}
			}

		}(s, i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	res := []string{} // I want to also initialize the slice
	for r := range ch {
		validResults[r.position] = r.word
	}

	if len(validResults) == 0 {
		return res
	}

	for i := range stringArray {
		if word, exists := validResults[i]; exists {
			res = append(res, word)
		}
	}
	return res
}
