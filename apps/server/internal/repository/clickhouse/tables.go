package clickhouse

import (
	"context"

	ch "github.com/ClickHouse/clickhouse-go/v2"
)

func (r *Repository) CreateTables(conn ch.Conn) error {

	query := `
	CREATE TABLE IF NOT EXISTS api_logs
	(
		timestamp DateTime64(3),
		user_id UInt64,
		method String,
		path String,
		status UInt16,
		duration_ms UInt32
	)
	ENGINE = MergeTree()
	ORDER BY (timestamp, user_id)
	`

	return r.conn.Exec(context.Background(), query)
}