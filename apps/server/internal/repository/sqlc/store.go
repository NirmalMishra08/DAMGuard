package sqlc

import (
	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

type Store struct {
	db    *pgxpool.Pool
	sqlDB *sql.DB

	queries *Queries
}

func NewStore(pool *pgxpool.Pool) *Store {
	sqlDB := stdlib.OpenDBFromPool(pool)

	return &Store{
		db:      pool,
		sqlDB:   sqlDB,
		queries: New(sqlDB),
	}
}
