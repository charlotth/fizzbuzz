package fizzbuzz

import (
	"io"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func String(opt ...Option) string {
	var sb strings.Builder
	if err := Write(&sb, opt...); err != nil {
		return ""
	}
	return sb.String()
}

// Write is our main entry point to compute the fizzbuzz string
func Write(w io.Writer, opt ...Option) error {
	if w == nil {
		return errors.New("nil writer")
	}

	// configure options
	options := newOptions(opt...)
	if options.From <= 0 {
		return errors.New("invalid from")
	}
	if options.To <= 0 {
		return errors.New("invalid to")
	}

	for i := options.From; i <= options.To; i++ {
		isNumber := true
		if i%options.Fizz.Multiple == 0 {
			if _, err := w.Write([]byte(options.Fizz.Str)); err != nil {
				return errors.Wrapf(err, "writing fizz: %d", i)
			}
			isNumber = false
		}
		if i%options.Buzz.Multiple == 0 {
			if _, err := w.Write([]byte(options.Buzz.Str)); err != nil {
				return errors.Wrapf(err, "writing buzz: %d", i)
			}
			isNumber = false
		}
		if isNumber {
			if _, err := w.Write([]byte(strconv.Itoa(i))); err != nil {
				return errors.Wrapf(err, "writing number: %d", i)
			}
		}

		if i != options.To {
			if _, err := w.Write([]byte(options.Separator)); err != nil {
				return errors.Wrapf(err, "writing separator: %d", i)
			}
		}
	}
	return nil
}
