package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	if hasBackslash(input) {
		return UnpackWithBackslash(input)
	}

	var builder strings.Builder
	inputLen := len(input)

	for key, currSymbol := range input {
		if key == 0 && !isEnglishLetter(currSymbol) {
			if unicode.IsLetter(currSymbol) || unicode.IsDigit(currSymbol) {
				return "", ErrInvalidString
			}
			return "", nil
		}
		if !unicode.IsDigit(currSymbol) && !isEnglishLetter(currSymbol) {
			continue
		}
		if key == inputLen-1 {
			builder.WriteRune(currSymbol)
			continue
		}

		_, curErr := strconv.Atoi(string(currSymbol))

		nextSymbol := input[key+1]
		nextDigit, nextErr := strconv.Atoi(string(nextSymbol))
		if curErr == nil {
			if nextErr == nil {
				return "", ErrInvalidString
			}
			continue
		}

		if nextErr == nil {
			builder.WriteString(strings.Repeat(string(currSymbol), nextDigit))
		} else {
			builder.WriteRune(currSymbol)
		}
	}

	return builder.String(), nil
}

func UnpackWithBackslash(input string) (string, error) {
	var builder strings.Builder
	inputLen := len(input)
	skipKeys := make(map[int]byte, inputLen)

	for key, currSymbol := range input {
		_, ok := skipKeys[key]
		if ok {
			continue
		}

		if currSymbol == 92 {
			if (inputLen - key - 1) >= 2 {
				nextSymbol := input[key+1]
				if nextSymbol == 110 {
					return "", ErrInvalidString
				}

				nextNextSymbol := input[key+2]
				if nextNextSymbol == 92 {
					skipKeys[key+2] = nextNextSymbol
					continue
				}

				_, ok := skipKeys[key+1]

				if ok {
					builder.WriteRune(currSymbol)
					continue
				}

				nextNextDigit, nextNextErr := strconv.Atoi(string(nextNextSymbol))
				if nextNextErr == nil {
					builder.WriteString(strings.Repeat(string(nextSymbol), nextNextDigit))
					skipKeys[key+1] = nextSymbol
					skipKeys[key+2] = nextNextSymbol
				}
			}
			continue
		}
		builder.WriteRune(currSymbol)
	}

	return builder.String(), nil
}

func isEnglishLetter(r rune) bool {
	if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
		return false
	}
	return true
}

func hasBackslash(input string) bool {
	hasBackslash := false
	for _, v := range input {
		if v == 92 {
			hasBackslash = true
			break
		}
	}
	return hasBackslash
}
