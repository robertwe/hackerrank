package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_handler(t *testing.T) {
	type args struct {
		a int
		b int
	}
	// tests := []struct {
	// 	name string
	// 	args args
	// 	want int
	// }{
	// 	{
	// 		"first",
	// 		args{1, 2},
	// 		3,
	// 	},
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := handler(tt.args.a, tt.args.b); got != tt.want {
	// 			t.Errorf("handler() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}
