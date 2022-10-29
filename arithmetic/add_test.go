package arithmetic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddInt(t *testing.T) {
	type args struct {
		x int64
		y int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Test 1",
			args: args{
				x: 12,
				y: 20,
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AddInt(tt.args.x, tt.args.y)
			assert.Equal(t, tt.want, got, fmt.Sprintf("AddInt() = %v, want %v", got, tt.want))
		})
	}
}

func TestAddFloat(t *testing.T) {
	type args struct {
		x, y float64
	}

	tests := []struct {
		name     string
		args     args
		expected float64
	}{
		{
			name: "float test 1",
			args: args{
				x: float64(12.40),
				y: float64(10.60),
			},
			expected: float64(23),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AddFloat(tt.args.x, tt.args.y)
			assert.Equal(t, tt.expected, got, fmt.Sprintf("AddInt() = %v, want %v", got, tt.expected))
		})
	}
}
