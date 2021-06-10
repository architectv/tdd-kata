package string_calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := map[string]struct{
		input string
		want  int
	}{
		"empty input": {
			input: "",
			want:  0,
		},
		"one number": {
			input: "2112",
			want:  2112,
		},
		"two numbers": {
			input: "2,3",
			want:  5,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Add(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}
