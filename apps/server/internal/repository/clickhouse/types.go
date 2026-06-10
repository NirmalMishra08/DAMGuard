package clickhouse

import "time"

type APILog struct {
	Timestamp  time.Time
	UserID     uint64
	Method     string
	Path       string
	Status     uint16
	DurationMS uint32
}