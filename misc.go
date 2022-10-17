package mystdlib

import (
	"errors"
	"math/rand"

	"github.com/google/uuid"
)

// RandomInt returns a random number between min and max
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// MustUUIDFromString returns UUID based on its string representation. Panics if `s` cannot be parsed! Deprecated, use uuid.MustParse.
func MustUUIDFromString(s string) uuid.UUID {
	u, err := uuid.Parse(s)
	if err != nil {
		panic(err)
	}
	return u
}

// PanicIfErr panics if e is not nil
func PanicIfErr(e error) {
	if e != nil {
		panic(e)
	}
}

// IsStringInSlice returns true if `needle` is in `haystack`
func IsStringInSlice(needle string, haystack []string) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}

	return false
}

func Trim(input string, maxLen int, trimSuffix string) (string, error) {
	inputLen := len(input)
	suffixLen := len(trimSuffix)

	if inputLen <= maxLen {
		return input, nil
	}

	if suffixLen > maxLen {
		return "", errors.New("trimSuffix length must not be longer than maxLen")
	}

	return input[:maxLen-suffixLen] + trimSuffix, nil
}

func Test() string {
	return "hello"
}
