package mystdlib

import (
	"math/rand"
)

// RandomInt returns a random number between min and max
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// PanicIfErr panics if e is not nil
func PanicIfErr(e error) {
	if e != nil {
		panic(e)
	}
}

// IsStringInSlice returns true if `needle` is in `haystack`
func IsStringInSlice(needle string, haystack []string) bool {
	return IsInSlice(needle, haystack)
}

// IsInSlice returns true if `needle` is in `haystack`. Implemented with generics.
func IsInSlice[T comparable](needle T, haystack []T) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}

	return false
}

// Will panic if `trimSuffix` is longer than `maxLen`.
func TrimString(input string, maxLen int, trimSuffix string) string {
	inputLen := len(input)
	suffixLen := len(trimSuffix)

	if inputLen <= maxLen {
		return input
	}

	if suffixLen > maxLen {
		panic("trimSuffix length must not be longer than maxLen")
	}

	return input[:maxLen-suffixLen] + trimSuffix
}
