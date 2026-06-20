package postgres

import (
	"strings"
)

func ExtractQuery(line string) (string, bool) {
	const marker = "LOG: statement:"
	if !strings.Contains(line, marker) {
		return "", false
	}

	parts := strings.Split(line, marker)

	if len(parts) < 2 {
		return "", false
	}

	query:= strings.TrimSpace(parts[1]);

	return query, true
}
