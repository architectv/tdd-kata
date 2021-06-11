package string_calculator_test

import (
	"fmt"
	"github.com/architectv/tdd-kata/string_calculator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
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
			input: "21",
			want:  21,
		},
		"two numbers": {
			input: "2,3",
			want:  5,
		},
		"some numbers": {
			input: "1,2,3,4,5",
			want:  15,
		},
		"(ok) new lines between numbers": {
			input: "1\n2,3",
			want:  6,
		},
		"(wrong) new lines between numbers": {
			input: "1,\n",
			want:  0,
			err:   fmt.Errorf("wrong number"),
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
		"negative & positive numbers": {
			input: "-1,1,2,3,-2,-3",
			want:  0,
			err:   fmt.Errorf("negatives not allowed: -1,-2,-3"),
		},
		"big numbers": {
			input: "1,2,3,1001,4,5",
			want:  15,
		},
		"some symbols in delimiter": {
			input: "//[***]\n1***2***3",
			want:  6,
		},
		"multiple delimiters with one char": {
			input: "//[*][%]\n1*2%3",
			want:  6,
		},
		"multiple delimiters with some chars": {
			input: "//[***][%][;;]\n1***2%3;;4***5",
			want:  15,
		},
		"wrong delimiters": {
			input: "//***][%][;;]\n1***2%3;;4***5",
			want:  0,
			err:   fmt.Errorf("wrong delimiter"),
		},
	}

	for name, tc := range tests {
		// Note: https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		// Also: https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := string_calculator.Add(tc.input)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
