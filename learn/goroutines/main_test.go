package main

import (
	"sort"
	"testing"
)

func TestParallelCount(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "count to 5",
			n:    5,
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "count to 1",
			n:    1,
			want: []int{1},
		},
		{
			name: "count to 10",
			n:    10,
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParallelCount(tt.n)

			// Sort the result since goroutines may complete in any order
			sort.Ints(got)

			// Check length
			if len(got) != len(tt.want) {
				t.Errorf("ParallelCount(%d) returned %d numbers, want %d", tt.n, len(got), len(tt.want))
			}

			// Check contents
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("ParallelCount(%d) = %v, want %v", tt.n, got, tt.want)
					break
				}
			}
		})
	}
}
