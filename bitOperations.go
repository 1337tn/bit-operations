package bitOperations

import (
	"strings"
)

const (
	UPPER = 1 << iota
	LOWER
	CAP
	REV
)

func processString(str string, conf byte) string {
	rev := func(s string) string {
		runes := []rune(s)
		n := len(runes)
		for i := 0; i < n/2; i++ {
			runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
		}
		return string(runes)
	}

	title := func(input string) string {
		words := strings.Fields(input)
		smallwords := " a an on the to "

		for index, word := range words {
			if strings.Contains(smallwords, " "+word+" ") {
				words[index] = word
			} else {
				words[index] = strings.Title(word)
			}
		}
		return strings.Join(words, " ")
	}

	if (conf & UPPER) != 0 {
		str = strings.ToUpper(str)
	}
	if (conf & LOWER) != 0 {
		str = strings.ToLower(str)
	}
	if (conf & CAP) != 0 {
		str = title(strings.ToLower(str))
	}
	if (conf & REV) != 0 {
		str = rev(str)
	}
	return str
}

func sameSign(a, b int) bool {
	return a^b >= 0
}

func setNthBit(num, position int) int {
	return num | (1 << (position - 1))
}

func unsetNthBit(num, position int) int {
	return num &^ (1 << (position - 1))
}

func getNthBit(num, position int) int {
	return num & (1 << (position - 1))
}