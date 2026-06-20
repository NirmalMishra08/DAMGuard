package postgres

import (
	"main/internal/repository/clickhouse"
	"time"

	"github.com/google/uuid"
)

func BuildAuditEvents(query string) *clickhouse.AuditEvent {
	return &clickhouse.AuditEvent{
		EventID:      uuid.New(),
		Timestamp:    time.Now(),
		DatabaseType: "postgres",
		Query:        query,
		QueryType:    DetectQueryType(query),
	}
}
