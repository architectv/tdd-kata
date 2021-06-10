package string_calculator

import (
	"strconv"
	"strings"
)

func Add(input string) int {
	numbers := strings.Split(input, ",")

	sum := 0
	for _, value := range numbers {
		number, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			sum += int(number)
		}
	}

	return sum
}
