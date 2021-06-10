package string_calculator_test

import (
	"github.com/architectv/tdd-kata/string_calculator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := map[string]struct {
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
			got := string_calculator.Add(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}
