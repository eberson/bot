package strs

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

//WhiteSpace is a strings that represents a white space
const WhiteSpace = " "

// Empty returns a empty string
func Empty() string {
	return ""
}

// IsEmpty verifies if a string is empty
func IsEmpty(v string) bool {
	return v == Empty()
}

//IsEmptyOrWhiteSpace verifies if a string is empty or is a white space
func IsEmptyOrWhiteSpace(v string) bool {
	return IsEmpty(strings.Trim(v, WhiteSpace))
}

// IsNotEmpty verifies if a string is not empty
func IsNotEmpty(v string) bool {
	return !IsEmpty(v)
}

// DefaultString returns a default string if the first option is empty
func DefaultString(v1, v2 string) string {
	if IsEmpty(v1) {
		return v2
	}
	return v1
}

// RightPad pads the input string with extra characters from the right side
func RightPad(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

// LeftPad pads the input string with extra characters from the left side
func LeftPad(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

//NormalizeText is a method to remove runes from text
func NormalizeText(text string) string {
	transformer := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	text, _, _ = transform.String(transformer, text)

	return text
}
