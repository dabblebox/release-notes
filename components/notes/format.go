package notes

import (
	"bytes"
	"fmt"
)

// Format ...
func Format(notes []string) string {
	var buffer bytes.Buffer

	for _, n := range notes {
		buffer.WriteString(fmt.Sprintf("\\n + %s", n))
	}

	return buffer.String()
}
