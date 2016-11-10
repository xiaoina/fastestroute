package Utilities

import (
	"strings"
)

var (
	replacer = strings.NewReplacer(
		" ", "",
		",", "",
		"-", "",
		".", "",
		"\n", "")
)

//TrimString removes unnessary data for searching with google's API
func TrimString(s string) string {
	return replacer.Replace(s)
}
