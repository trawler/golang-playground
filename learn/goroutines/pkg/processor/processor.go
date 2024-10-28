/*
Your task is to implement ParallelVowelCount function that:

Processes each string in parallel using goroutines
Counts vowels in each string
Returns the counts in the same order as input strings
Uses proper concurrency patterns we learned (channels, WaitGroup, defer)
*/

package processor

import (
	"strings"
	"sync"
)

var vowels = "aeiou"

type result struct {
	position int
	count    int
}

func ParallelVowelCount(stringArray []string) []int {
	var wg sync.WaitGroup
	ch := make(chan result, len(stringArray)) // buffered channel

	for i, s := range stringArray {
		wg.Add(1)
		go func(str string, pos int) {
			defer wg.Done()

			ch <- result{
				position: pos,
				count:    getNumberOfVowels(str),
			}
		}(s, i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	res := make([]int, len(stringArray))
	for r := range ch {
		res[r.position] = r.count
	}

	return res
}

func getNumberOfVowels(s string) int {
	var count int
	for _, r := range s {
		if strings.Contains(vowels, string(r)) {
			count++
		}
	}
	return count
}
