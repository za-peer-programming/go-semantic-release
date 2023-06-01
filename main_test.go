package main

import "testing"

func Test_message(t *testing.T) {
	type args struct {
		who []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Basic Test",
			args: args{
				who: []string{"there"},
			},
			want: "Hello there",
		},
		{
			name: "Basic Test no args",
			want: "Hello there!",
		},
		{
			name: "Basic Args Array",
			args: args{
				who: []string{"1", "2"},
			},
			want: "Hello 1, 2",
		},
		{
			name: "Basic 3 Args",
			args: args{
				who: []string{"1", "2", "No"},
			},
			want: "Hello 1, 2, No",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Message(tt.args.who...); got != tt.want {
				t.Errorf("message() = %v, want %v", got, tt.want)
			}
		})
	}
}
