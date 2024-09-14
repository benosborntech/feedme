package utils

import "strings"

func GenerateKey(elems ...string) string {
	return strings.Join(elems, "-")
}
