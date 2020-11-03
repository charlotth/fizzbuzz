package fizzbuzz_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tcharlot-datasweet/fizzbuzz/pkg/fizzbuzz"
)

func TestDefault(t *testing.T) {
	// Fizzbuzz from 1 to 100
	assert.Equal(t, "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,17,fizz,19,buzz,fizz,22,23,fizz,buzz,26,fizz,28,29,fizzbuzz,31,32,fizz,34,buzz,fizz,37,38,fizz,buzz,41,fizz,43,44,fizzbuzz,46,47,fizz,49,buzz,fizz,52,53,fizz,buzz,56,fizz,58,59,fizzbuzz,61,62,fizz,64,buzz,fizz,67,68,fizz,buzz,71,fizz,73,74,fizzbuzz,76,77,fizz,79,buzz,fizz,82,83,fizz,buzz,86,fizz,88,89,fizzbuzz,91,92,fizz,94,buzz,fizz,97,98,fizz,buzz", fizzbuzz.String())
}

func TestFromOption(t *testing.T) {
	assert.Equal(t, "fizzbuzz,91,92,fizz,94,buzz,fizz,97,98,fizz,buzz", fizzbuzz.String(fizzbuzz.From(90)))
	assert.Equal(t, "", fizzbuzz.String(fizzbuzz.From(-1)))
	assert.Equal(t, "", fizzbuzz.String(fizzbuzz.From(0)))
	assert.Equal(t, "buzz", fizzbuzz.String(fizzbuzz.From(100)))
	assert.Equal(t, "", fizzbuzz.String(fizzbuzz.From(101)))
}

func TestToOption(t *testing.T) {
	assert.Equal(t, "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16", fizzbuzz.String(fizzbuzz.To(16)))
	assert.Equal(t, "", fizzbuzz.String(fizzbuzz.To(-1)))
	assert.Equal(t, "", fizzbuzz.String(fizzbuzz.To(0)))
}

func TestSeparatorOption(t *testing.T) {
	assert.Equal(t, "fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz 16", fizzbuzz.String(
		fizzbuzz.From(3),
		fizzbuzz.To(16),
		fizzbuzz.Separator(" "),
	))
}

func TestFizzOption(t *testing.T) {
	assert.Equal(t, "1,2,FIZZZZZ,4,buzz,FIZZZZZ,7,8,FIZZZZZ,buzz,11,FIZZZZZ,13,14,FIZZZZZbuzz,16", fizzbuzz.String(
		fizzbuzz.To(16),
		fizzbuzz.Fizz(3, "FIZZZZZ"),
	))
	assert.Equal(t, "1,2,3,4,buzz,6,fizz,8,9,buzz,11,12,13,fizz,buzz,16", fizzbuzz.String(
		fizzbuzz.To(16),
		fizzbuzz.Fizz(7, "fizz"),
	))
	assert.Equal(t, "1,2,3,4,fizzbuzz,6,7,8,9,fizzbuzz,11,12,13,14,fizzbuzz,16", fizzbuzz.String(
		fizzbuzz.To(16),
		fizzbuzz.Fizz(5, "fizz"),
	))
}

func TestBuzzOption(t *testing.T) {
	assert.Equal(t, "1,2,fizz,4,BUZZZZZZ,fizz,7,8,fizz,BUZZZZZZ,11,fizz,13,14,fizzBUZZZZZZ,16", fizzbuzz.String(
		fizzbuzz.To(16),
		fizzbuzz.Buzz(5, "BUZZZZZZ"),
	))
	assert.Equal(t, "1,2,fizz,4,5,fizzbuzz,7,8,fizz,10,11,fizzbuzz,13,14,fizz,16", fizzbuzz.String(
		fizzbuzz.To(16),
		fizzbuzz.Buzz(6, "buzz"),
	))
	assert.Equal(t, "1,2,fizzbuzz,4,5,fizzbuzz,7,8,fizzbuzz,10,11,fizzbuzz,13,14,fizzbuzz,16", fizzbuzz.String(
		fizzbuzz.To(16),
		fizzbuzz.Buzz(3, "buzz"),
	))
}
