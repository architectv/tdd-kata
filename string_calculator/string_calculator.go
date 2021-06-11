package string_calculator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func NewDelimiters(line string) ([]string, error) {
	var delimiters []string

	line = strings.TrimPrefix(line, "//")
	if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
		re := regexp.MustCompile(`\[(.*?)\]`)
		delimiters = re.FindAllString(line, -1)
		for i := range delimiters {
			delimiters[i] = strings.TrimPrefix(delimiters[i], "[")
			delimiters[i] = strings.TrimSuffix(delimiters[i], "]")
		}
	} else if utf8.RuneCountInString(line) == 1 {
		delimiters = append(delimiters, line)
	} else {
		return nil, fmt.Errorf("wrong delimiter")
	}

	return delimiters, nil
}

func GetDelimiters(input string) (string, []string, error) {
	var delimiters []string

	// Change delimiter: `//[delimiter]\n[numbers...]`
	if strings.HasPrefix(input, "//") {
		lines := strings.SplitN(input, "\n", 2)
		if len(lines) != 2 {
			return "", nil, fmt.Errorf("wrong delimiter")
		}

		var err error
		delimiters, err = NewDelimiters(lines[0])
		if err != nil {
			return "", nil, err
		}

		input = lines[1]
	} else {
		delimiters = append(delimiters, ",")
	}

	delimiters = append(delimiters, "\n")

	return input, delimiters, nil
}

func Add(input string) (int, error) {
	if input == "" {
		return 0, nil
	}

	numbersString, delimiters, err := GetDelimiters(input)
	if err != nil {
		return 0, err
	}

	for i := range delimiters {
		delimiters[i] = regexp.QuoteMeta(delimiters[i])
	}

	re := regexp.MustCompile(strings.Join(delimiters, "|"))
	numbers := re.Split(numbersString, -1)

	sum := 0
	var negatives []string
	negativeFlag := false

	for _, value := range numbers {
		number, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("wrong number")
		}

		if number < 0 {
			negatives = append(negatives, value)
			negativeFlag = true
		}

		if !negativeFlag && number <= 1000 {
			sum += int(number)
		}
	}

	if len(negatives) != 0 {
		errString := "negatives not allowed: " + strings.Join(negatives, ",")
		return 0, fmt.Errorf(errString)
	}

	return sum, nil
}
