package processor

import (
	"reflect"
	"testing"
)

func TestParallelVowelCount(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		output []int
	}{
		{
			name:   "basic words",
			input:  []string{"hello", "world", "go"},
			output: []int{2, 1, 1},
		},
		{
			name:   "empty strings",
			input:  []string{"", "", "a"},
			output: []int{0, 0, 1},
		},
		{
			name:   "longer words",
			input:  []string{"beautiful", "exercise", "automation"},
			output: []int{5, 4, 6},
		},
		{
			name:   "single string",
			input:  []string{"aeiou"},
			output: []int{5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParallelVowelCount(tt.input)
			if !reflect.DeepEqual(got, tt.output) {
				t.Errorf("ParallelVowelCount(%v) = %v, want %v", tt.input, got, tt.output)
			}
		})
	}
}
