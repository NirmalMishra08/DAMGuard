package clickhouse

import (
	"context"
	"time"

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

func (r *Repository) InsertQueryEvent(event AuditEvent) error {
	ctx := context.Background()

	batch, err := r.conn.PrepareBatch(ctx, `
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
	`)
	if err != nil {
		return err
	}

	if err := batch.Append(
		event.EventID,
		event.Timestamp,
		event.DatabaseID,
		event.DatabaseName,
		event.DatabaseType,
		event.Username,
		event.Query,
		event.QueryType,
		event.ClientIP,
	); err != nil {
		return err
	}

	return batch.Send()
}
