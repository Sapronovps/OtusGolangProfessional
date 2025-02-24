package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	if hasBackslash(input) {
		return UnpackWithBackslash(input)
	}

	var builder strings.Builder
	inputRunes := []rune(input)
	inputLen := len(inputRunes)
	skipKeys := make(map[int]byte, inputLen)

	i := -1
	for _, currSymbol := range input {
		i++
		// Если символ есть в карте для пропуска
		_, isSkip := skipKeys[i]
		if isSkip {
			continue
		}

		// Если первый символ цифра - тогда ошибка
		_, isCurrSymbolDigit := strconv.Atoi(string(currSymbol))
		if i == 0 && isCurrSymbolDigit == nil {
			return "", ErrInvalidString
		}

		if (inputLen - i) > 1 {
			nextSymbol := inputRunes[i+1]
			nextSymbolDigit, isNextSymbolDigit := strconv.Atoi(string(nextSymbol))

			if (inputLen - i) > 2 {
				nextNextSymbol := inputRunes[i+2]
				_, isNNextDigit := strconv.Atoi(string(nextNextSymbol))
				if isNextSymbolDigit == nil && isNNextDigit == nil || isCurrSymbolDigit == nil && isNextSymbolDigit == nil {
					return "", ErrInvalidString
				}
			}

			// Если текущий символ не цифра, а следующий символ цифра, тогда делаем repeat символа
			if isCurrSymbolDigit != nil && isNextSymbolDigit == nil {
				builder.WriteString(strings.Repeat(string(currSymbol), nextSymbolDigit))
				skipKeys[i+1] = byte(nextSymbol)
				continue
			}
		}

		builder.WriteRune(currSymbol)
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

		if currSymbol != 92 {
			builder.WriteRune(currSymbol)
			continue
		}

		if (inputLen - key - 1) < 2 {
			continue
		}

		nextSymbol := input[key+1]
		if nextSymbol == 110 {
			return "", ErrInvalidString
		}

		nextNextSymbol := input[key+2]
		if nextNextSymbol == 92 {
			skipKeys[key+2] = nextNextSymbol
			continue
		}

		_, nextSymbolSkip := skipKeys[key+1]

		if nextSymbolSkip {
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

	return builder.String(), nil
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
