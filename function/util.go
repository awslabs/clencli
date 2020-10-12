package function

import (
	"strings"
)

// GetStringInBetweenTwoString return the string between two strings
func GetStringInBetweenTwoString(str string, startS string, endS string) (result string, found bool) {
	s := strings.Index(str, startS)
	if s == -1 {
		return result, false
	}
	newS := str[s+len(startS):]
	e := strings.Index(newS, endS)
	if e == -1 {
		return result, false
	}
	result = newS[:e]
	return result, true
}

// GetStringBetweenDoubleQuotes return the string between double quotes
func GetStringBetweenDoubleQuotes(str string) (result string, found bool) {
	return GetStringInBetweenTwoString(str, "\"", "\"")
}


// GetStringTrimmed splits the string by the given separator and trims it by removing all spaces in between
func GetStringTrimmed(s string, sep string) []string {
	slc := strings.Split(s, sep)
	for i := range slc {
		slc[i] = strings.TrimSpace(slc[i])
	}
	return slc
}