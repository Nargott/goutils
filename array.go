package goutils

import "strings"

/**
Returns true if given string is in array
 */
func IsStringInArray(str string, arr []string) bool {
	for _, record := range arr {
		if strings.EqualFold(str, record) {
			return true
		}
	}

	return false
}
