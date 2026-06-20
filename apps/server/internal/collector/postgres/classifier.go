package postgres

import "strings"

func DetectQueryType(query string) string {
	q := strings.TrimSpace(strings.ToUpper(query))

	switch {
	case strings.HasPrefix(q, "SELECT"):
		return "SELECT"

	case strings.HasPrefix(q, "INSERT"):
		return "INSERT"

	case strings.HasPrefix(q, "UPDATE"):
		return "UPDATE"

	case strings.HasPrefix(q, "DELETE"):
		return "DELETE"

	case strings.HasPrefix(q, "DROP"):
		return "DROP"

	case strings.HasPrefix(q, "CREATE"):
		return "CREATE"

	case strings.HasPrefix(q, "ALTER"):
		return "ALTER"

	case strings.HasPrefix(q, "TRUNCATE"):
		return "TRUNCATE"

	case strings.HasPrefix(q, "GRANT"):
		return "GRANT"

	default:
		return "UNKNOWN"
	}
}