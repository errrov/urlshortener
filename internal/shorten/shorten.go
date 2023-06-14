package shorten

import (
	"net/url"
	"strings"
)

const alphabet = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM0123456789_"

var alphabetLen = uint32(len(alphabet))

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
	Reverse(digits)
	for _, digit := range digits {
		shortString.WriteString(string(alphabet[digit]))
	}
	return shortString.String()
}

func Reverse(nums []uint32) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func CreateShortURL(burl, id string) (string, error) {
	parse, err := url.Parse(burl)
	if err != nil {
		return "", err
	}

	parse.Path = id
	return parse.String(), nil
}
