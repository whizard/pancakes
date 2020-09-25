package ihop

import (
	"bytes"
	"fmt"
	"regexp"
)

const (
	ValidStackChars      = `^[-\+]*$`
	HappyChar       byte = '+'
	PlainChar       byte = '-'
)

// PancakeFlips Given a stack of pancakes represented by a string HappyChar or PlainChar characters
// calculate the minimum number of flips to have all HappyChar characters
func PancakeFlips(stack string) int {
	flips := 0
	str := []byte(stack)

	for {
		pos := bytes.LastIndex(str, []byte(string(PlainChar)))
		if pos == -1 {
			break
		} else {
			flips++
			for i := 0; i <= pos; i++ {
				if str[i] == PlainChar {
					str[i] = HappyChar
				} else {
					str[i] = PlainChar
				}
			}
		}
	}
	return flips
}

// ValidateStack Validate that a stack string contains only HappyChar or PlainChar characters
func ValidateStack(stack string) error {
	rgx := regexp.MustCompile(ValidStackChars)
	if !rgx.MatchString(stack) {
		return fmt.Errorf("Stack of pancakes string contains invalid characters ('%s'). Valid characters are '-' or '+'.", stack)
	}
	return nil
}
