package models

import "time"

type AuditEvent struct {
	EventID   string
	Timestamp time.Time

	DatabaseID   string
	DatabaseName string
	DatabaseType string

	Username string
	ClientIP string

	Query     string
	QueryType string
}
