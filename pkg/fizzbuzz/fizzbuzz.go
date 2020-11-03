package fizzbuzz

import (
	"strconv"
	"strings"
)

// String is our main entry point to output a fizzbuzz string
func String(opt ...Option) string {
	// configure options
	options := newOptions(opt...)

	// Check options
	if options.From <= 0 || options.To <= 0 {
		return ""
	}

	var sb strings.Builder

	for i := options.From; i <= options.To; i++ {
		isNumber := true
		if i%options.Fizz.Multiple == 0 {
			sb.WriteString(options.Fizz.Str)
			isNumber = false
		}
		if i%options.Buzz.Multiple == 0 {
			sb.WriteString(options.Buzz.Str)
			isNumber = false
		}
		if isNumber {
			sb.WriteString(strconv.Itoa(i))
		}
		sb.WriteString(options.Separator)
	}

	return strings.TrimSuffix(sb.String(), options.Separator)
}
