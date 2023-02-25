package utils

import (
	"regexp"
	"strings"

	"github.com/essentialkaos/translit/v2"
)

func Slugificator(s string) string {
	return strings.ReplaceAll(
		strings.ToLower(
			regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(
				translit.EncodeToPCGN(s), "",
			),
		),
		" ", "-",
	)
}
