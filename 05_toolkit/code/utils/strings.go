package utils

import "strings"

//transform a sentence to all caps an exclamation point
func MakeExcited(sentence string) string{
	return strings.ToUpper(sentence) + "!"
}