package links

import (
	"regexp"
	"strings"
)

// URL ...
type URL struct {
	ReplaceRegEx  string
	ReplaceToken  string
	TokenizedLink string
}

// Insert ...
func Insert(url URL, comment string) string {
	var re = regexp.MustCompile(url.ReplaceRegEx)

	for _, match := range re.FindAllString(comment, -1) {
		url.TokenizedLink = strings.Replace(url.TokenizedLink, url.ReplaceToken, match, -1)
		comment = strings.Replace(comment, match, url.TokenizedLink, -1)
	}

	return comment
}
