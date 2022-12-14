package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const CharBackslash = 92

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputString string) (string, error) {
	var b strings.Builder

	runes := []rune(inputString)
	for position := 0; position < len(runes); position++ {
		if unicode.IsDigit(runes[position]) {
			return inputString, ErrInvalidString
		}

		if CharBackslash == runes[position] {
			if len(runes) > position+1 &&
				(unicode.IsDigit(runes[position+1]) || CharBackslash == runes[position+1]) {
				position++
			} else {
				return inputString, ErrInvalidString
			}
		}

		if len(runes) > position+1 && unicode.IsDigit(runes[position+1]) {
			count, _ := strconv.Atoi(string(runes[position+1]))
			if count > 0 {
				b.WriteString(strings.Repeat(string(runes[position]), count))
			}
			position++
		} else {
			b.WriteRune(runes[position])
		}
	}

	return b.String(), nil
}
