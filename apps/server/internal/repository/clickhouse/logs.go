package clickhouse

import (
	"context"
	"time"

	ch "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/google/uuid"
)

type AuditEvent struct {
	EventID      uuid.UUID
	Timestamp    time.Time
	DatabaseID   string
	DatabaseName string
	DatabaseType string
	Username     string
	Query        string
	QueryType    string
	ClientIP     string
}

func (r *Repository) InsertAuditEvent(
	conn ch.Conn,
	event AuditEvent,
) error {

	return conn.Exec(
		context.Background(),
		`
		INSERT INTO query_events
		(
			event_id,
			timestamp,
			database_id,
			database_name,
			database_type,
			username,
			query,
			query_type,
			client_ip
		)
		VALUES (toUUID(?), ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		event.EventID,
		event.Timestamp,
		event.DatabaseID,
		event.DatabaseName,
		event.DatabaseType,
		event.Username,
		event.Query,
		event.QueryType,
		event.ClientIP,
	)
}
