package shorten

import (
	"strings"
)

const alphabet = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM0123456789_"

var alphabetLen = uint32(len(alphabet))

// we convert to base 63 -> 0_o. Extra comment should be removed later.

func Shorten(id uint32) string {
	var (
		digits      []uint32
		num         = id
		shortString strings.Builder
	)
	for num > 0 {
		digits = append(digits, num%alphabetLen)
		num /= alphabetLen
	}
	reverse(digits)
	for _, digit := range digits {
		shortString.WriteString(string(alphabet[digit]))
	}
	return shortString.String()
}

func reverse(nums []uint32) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
