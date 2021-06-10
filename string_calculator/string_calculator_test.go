package string_calculator_test

import (
	"github.com/architectv/tdd-kata/string_calculator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := map[string]struct {
		input   string
		want    int
		wantErr bool
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
		"(ok) new lines between numbers": {
			input: "1\n2,3",
			want:  6,
		},
		"(wrong) new lines between numbers": {
			input: "1,\n",
			want:  1,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := string_calculator.Add(tc.input)
			if tc.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tc.want, got)
		})
	}
}
