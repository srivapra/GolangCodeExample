package main

import (
	"reflect"
	"testing"
)

func TestSortArrayOf012(t *testing.T) {

	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "Test Case 1",
			input: []int{0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1},
			want:  []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2},
		},
		{
			name:  "Test Case 2",
			input: []int{0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1},
			want:  []int{0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 2, 2},
		},
		{
			name:  "Test Case 3",
			input: []int{},
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SortArrayOf012(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v ", got, tt.want)
			}
		})
	}

}
