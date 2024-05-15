package main

import (
	"testing"
)

func TestReverse(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Test case 1",
			input: "hellow",
			want:  "wolleh",
		},
		{
			name:  "Test case 2",
			input: "prashant",
			want:  "tnahsarp",
		},
		{
			name:  "Test case 3",
			input: "ram",
			want:  "tnahsarp",
		},
		{
			name:  "Test case 4",
			input: "",
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse(tt.input)
			if got != tt.want {
				t.Errorf("got %s want %s ", got, tt.want)
			}
		})
	}

}
