package string_calculator_test

import (
	"fmt"
	"github.com/architectv/tdd-kata/string_calculator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
		err   error
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
		"some numbers": {
			input: "1,2,3,4,5",
			want:  15,
		},
		"new lines between numbers [1]": {
			input: "1\n2,3",
			want:  6,
		},
		"new lines between numbers [2]": {
			input: "1,\n",
			want:  1,
		},
		"change delimiter": {
			input: "//;\n1;2",
			want:  3,
		},
		"negative number": {
			input: "-1",
			want:  0,
			err:   fmt.Errorf("negatives not allowed: -1"),
		},
		"negative numbers": {
			input: "-1,-2,-3",
			want:  0,
			err:   fmt.Errorf("negatives not allowed: -1,-2,-3"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := string_calculator.Add(tc.input)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
