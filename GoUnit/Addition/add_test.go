package main

import "testing"

// func TestAdd(t *testing.T) {
// 	res := Add(10, 2)
// 	if res == 12 {
// 		t.Logf("Result is correct got %d, want %d ", res, 30)
// 	} else {
// 		t.Errorf("Result is incorrect got %d, want %d ", res, 30)
// 	}
// }

func TestAdd(t *testing.T) {

	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{10, 20},
			want: 30,
		},
		{
			name: "Test Case 2",
			args: args{10, 2},
			want: 12,
		},
		{
			name: "Test Case 3",
			args: args{10, 5},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
