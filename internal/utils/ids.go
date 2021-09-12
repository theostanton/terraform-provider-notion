package utils

import "strings"

func NormalizeId(id string) string {
	return strings.ReplaceAll(id, "-", "")
}
