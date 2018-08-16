package iutil

import (
	"strings"
	"unicode"
)

// Strpos strpos()
func Strpos(haystack, needle string, offset int) int {
	return strings.Index(haystack, needle)
}

// stripos()
func Stripos(haystack, needle string, offset int) int {
	haystack = strings.ToLower(haystack)
	needle = strings.ToLower(needle)
	return strings.Index(haystack, needle)
}

// strrpos()
func Strrpos(haystack, needle string, offset int) int {
	return strings.LastIndex(haystack, needle)
}

// strripos()
func Strripos(haystack, needle string, offset int) int {
	haystack = strings.ToLower(haystack)
	needle = strings.ToLower(needle)
	return strings.LastIndex(haystack, needle)
}

// str_replace()
func StrReplace(search, replace, subject string, count int) string {
	return strings.Replace(subject, search, replace, count)
}

// strtoupper()
func Strtoupper(str string) string {
	return strings.ToUpper(str)
}

// strtolower()
func Strtolower(str string) string {
	return strings.ToLower(str)
}

// ucfirst()
func Ucfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return u + str[len(u):]
	}
	return ""
}

// lcfirst()
func Lcfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToLower(v))
		return u + str[len(u):]
	}
	return ""
}

// ucwords()
func Ucwords(str string) string {
	return strings.Title(str)
}

// substr()
func Substr(str string, start uint, length int) string {
	if start < 0 || length < -1 {
		return str
	}
	switch {
	case length == -1:
		return str[start:]
	case length == 0:
		return ""
	}
	end := int(start) + length
	if end > len(str) {
		end = len(str)
	}
	return str[start:end]
}

// strrev()
func Strrev(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
