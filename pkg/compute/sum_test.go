package compute

import (
	"testing"
)

func TestSum(t *testing.T) {
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
			name: "Sum 3 and 5",
			args: args{a: 3, b: 5},
			want: 8,
		},
		{
			name: "Sum 6 and 8",
			args: args{a: 6, b: 8},
			want: 14,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumNumber(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
