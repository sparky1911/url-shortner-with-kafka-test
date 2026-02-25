package utils

import (
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Encode(number uint64) string {
	if number == 0 {
		return string(alphabet[0])
	}

	var sb strings.Builder
	base := uint64(len(alphabet))

	for number > 0 {
		sb.WriteByte(alphabet[number%base])
		number /= base
	}

	return reverse(sb.String())
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
