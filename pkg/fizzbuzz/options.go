package fizzbuzz

// Options will configure the fizzbuzz
type Options struct {
	From      int
	To        int
	Fizz      Multiple
	Buzz      Multiple
	Separator string
}

// Multiple will replace multiple of {Multiple} by {Str}
type Multiple struct {
	Multiple int
	Str      string
}

type Option func(*Options)

func newOptions(opt ...Option) Options {
	options := Options{
		From:      1,
		To:        100,
		Fizz:      Multiple{3, "fizz"},
		Buzz:      Multiple{5, "buzz"},
		Separator: ",",
	}
	for _, o := range opt {
		o(&options)
	}
	return options
}

// From will starts the fizzbuzz to {v}
// Default to 1
func From(v int) Option {
	return func(opts *Options) {
		opts.From = v
	}
}

// To will ends the fizzbuzz to {v}
// Default to 100
func To(v int) Option {
	return func(opts *Options) {
		opts.To = v
	}
}

// Fizz will replace multiple of {value} by {str} value
// Default to 3, "fizz"
func Fizz(value int, str string) Option {
	return func(opts *Options) {
		opts.Fizz = Multiple{value, str}
	}
}

// Buzz will replace multiple of {value} by {str} value
// Default to 5, "buzz"
func Buzz(value int, str string) Option {
	return func(opts *Options) {
		opts.Buzz = Multiple{value, str}
	}
}

// Separator will sets the separator.
// Default to ","
func Separator(v string) Option {
	return func(opts *Options) {
		opts.Separator = v
	}
}
