package utils

import "strings"

func SlicesContainsString(list []string, str string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}
	return false
}

func SlicesLikeString(list []string, str string) bool {
	for _, s := range list {
		if strings.Contains(str, s) {
			return true
		}
	}
	return false
}
