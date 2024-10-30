package filter

import (
	"reflect"
	"testing"
)

func TestParallelWordFilter(t *testing.T) {
	tests := []struct {
		name      string
		words     []string
		minLength int
		want      []string
	}{
		{
			name:      "basic filter",
			words:     []string{"cat", "dog", "elephant", "bird", "rhinoceros", "ant"},
			minLength: 5,
			want:      []string{"elephant", "rhinoceros"},
		},
		{
			name:      "empty slice",
			words:     []string{},
			minLength: 3,
			want:      []string{},
		},
		{
			name:      "no matches",
			words:     []string{"a", "bb", "ccc"},
			minLength: 4,
			want:      []string{},
		},
		{
			name:      "all matches",
			words:     []string{"three", "four", "five"},
			minLength: 4,
			want:      []string{"three", "four", "five"},
		},
		{
			name:      "zero length",
			words:     []string{"", "a", "ab", "abc"},
			minLength: 0,
			want:      []string{"", "a", "ab", "abc"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParallelWordFilter(tt.words, tt.minLength)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParallelWordFilter(%v, %v) = %v, want %v",
					tt.words, tt.minLength, got, tt.want)
			}
		})
	}
}
