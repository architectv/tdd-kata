package string_calculator

import (
	"strconv"
	"strings"
)

func Add(input string) (int, error) {
	numbers := strings.FieldsFunc(input, func(c rune) bool {
		return c == ',' || c == '\n'
	})

	sum := 0
	for _, value := range numbers {
		number, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return 0, err
		}
		sum += int(number)
	}

	return sum, nil
}
