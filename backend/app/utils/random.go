package utils

import "crypto/rand"

const (
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"
	numbers      = "0123456789"
	alphaLetters = upperLetters + lowerLetters
	alphaNumbers = alphaLetters + numbers
)

type RandomOpts struct {
	Length  int
	Letters string
}

type Option func(*RandomOpts)

func Letters(v string) Option {
	return func(o *RandomOpts) {
		o.Letters = v
	}
}

func UpperCases() Option {
	return Letters(upperLetters)
}

func LowerCases() Option {
	return Letters(lowerLetters)
}

func Numbers() Option {
	return Letters(numbers)
}

func Length(v int) Option {
	return func(o *RandomOpts) {
		o.Length = v
	}
}

func GetRandomString(opts ...Option) (string, error) {
	o := &RandomOpts{
		Length:  16,
		Letters: alphaNumbers,
	}

	for _, opt := range opts {
		opt(o)
	}

	b := make([]byte, o.Length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	result := ""

	for _, v := range b {
		result += string(o.Letters[int(v)%len(o.Letters)])
	}

	return result, nil
}
