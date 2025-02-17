package shorten

import (
	"strings"

	"github.com/khapaev/solid-octo-meme/internal/utils"
)

var (
	charset    = []rune("ПЦ7ЛЙ4ЁЫЖЗЪ1ГЧФЭ8ЬЩКН0МВЕШ3ЯЮИТДБ9ОУРСА5Хпцлй2ёыжзъгчфэьщкнмвеш6яюитдбоурсах")
	charsetLen = uint32(len(charset))
)

func Shorten(number uint32) string {
	if number == 0 {
		return string(charset[0])
	}

	var indices []uint32
	for number > 0 {
		indices = append(indices, number%charsetLen)
		number /= charsetLen
	}

	utils.ReverseSlice(indices)

	var result strings.Builder
	for _, index := range indices {
		result.WriteRune(charset[index])
	}

	return result.String()
}
