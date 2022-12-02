package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputString string) (string, error) {

	var b strings.Builder

	for position := 0; position < len(inputString); position++ {
		if unicode.IsDigit(rune(inputString[position])) {
			return inputString, ErrInvalidString
		}

		if inputString[position] == 92 {
			if len(inputString) > position+1 && (unicode.IsDigit(rune(inputString[position+1])) || inputString[position+1] == 92) {
				position++
			} else {
				return inputString, ErrInvalidString
			}
		}

		if len(inputString) > position+1 && unicode.IsDigit(rune(inputString[position+1])) {
			count, _ := strconv.Atoi(string(inputString[position+1]))
			if count > 0 {
				b.WriteString(strings.Repeat(string(inputString[position]), count))
			}
			position++
		} else {
			b.WriteByte(inputString[position])
		}
	}

	return b.String(), nil
}
