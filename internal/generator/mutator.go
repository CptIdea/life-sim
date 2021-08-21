package generator

import (
	"strings"
)

func Mutate(code string) string {
	return strings.Replace(code, generateChar(), generateChar(), 1)
}
