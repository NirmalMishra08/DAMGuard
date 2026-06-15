package clickhouse

import (
	"context"

	ch "github.com/ClickHouse/clickhouse-go/v2"
)

func (r *Repository) CreateTables(conn ch.Conn) error {
	ctx := context.Background()

	// 2. Create query_events table
	queryEventsQuery := `
	CREATE TABLE IF NOT EXISTS query_events
	(
		event_id UUID,
		timestamp DateTime,
		database_id String,
		database_name String,
		database_type String,
		username String,
		query String,
		query_type String,
		client_ip String
	)
	ENGINE = MergeTree()
	ORDER BY timestamp
	` 
	if err := r.conn.Exec(ctx, queryEventsQuery); err != nil {
		return err
	}

	return nil
}


