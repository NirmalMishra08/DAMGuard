package clickhouse

import (
	"context"

	ch "github.com/ClickHouse/clickhouse-go/v2"
)

func(r *Repository) InsertAPILog(
	conn ch.Conn,
	log APILog,
) error {

	return r.conn.Exec(
		context.Background(),
		`
		INSERT INTO api_logs
		(
			timestamp,
			user_id,
			method,
			path,
			status,
			duration_ms
		)
		VALUES (?, ?, ?, ?, ?, ?)
		`,
		log.Timestamp,
		log.UserID,
		log.Method,
		log.Path,
		log.Status,
		log.DurationMS,
	)
}