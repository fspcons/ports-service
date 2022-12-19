package utils

import (
	"strings"
)

// IsEmpty checks whether a string after space trimming is empty or not.
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
