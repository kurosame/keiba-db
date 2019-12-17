package service

import (
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

// SanitizeHTML is sanitize HTML texts without policy
func SanitizeHTML(html string) []string {
	sh := bluemonday.NewPolicy().Sanitize(html)

	return strings.Split(sh, "\n")
}
