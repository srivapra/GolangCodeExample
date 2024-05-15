package main

import "testing"

func TestFind(t *testing.T) {
	arr := []int{12, 89, 23, 56, 33}
	type args struct {
		arr  []int
		item int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{arr, 23},
			want: true,
		},
		{
			name: "Test Cae 2",
			args: args{arr, 89},
			want: true,
		},
		{
			name: "Test Case 3",
			args: args{arr, 3},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Find(tt.args.arr, tt.args.item)
			if got != tt.want {
				t.Errorf("got %t want %t ", got, tt.want)

			}
		})
	}

}
