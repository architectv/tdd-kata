package string_calculator

import (
	"fmt"
	"strconv"
	"strings"
)

const defaultDelimiter = ','

func Add(input string) (int, error) {
	delimiter := defaultDelimiter
	runes := []rune(input)
	// Change delimiter: `//[delimiter]\n[numbers...]`
	if strings.HasPrefix(input, "//") && runes[3] == '\n' && len(runes) > 3 {
		delimiter = runes[2]
		input = string(runes[4:])
	}

	numbers := strings.FieldsFunc(input, func(c rune) bool {
		return c == delimiter || c == '\n'
	})

	sum := 0
	var negatives []string
	negativeFlag := false

	for _, value := range numbers {
		number, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return 0, err
		}

		if number < 0 {
			negatives = append(negatives, value)
			negativeFlag = true
		}

		if !negativeFlag {
			sum += int(number)
		}
	}

	if len(negatives) != 0 {
		errString := "negatives not allowed: " + strings.Join(negatives, ",")
		return 0, fmt.Errorf(errString)
	}

	return sum, nil
}
