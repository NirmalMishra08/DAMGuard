package clickhouse

import (
	"context"
	"fmt"
	"main/internal/config"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	ch "github.com/ClickHouse/clickhouse-go/v2"
)

type Repository struct {
	conn ch.Conn
}

func NewRepository(conn ch.Conn) *Repository {
	return &Repository{
		conn: conn,
	}
}

func New(cfg *config.Config) (clickhouse.Conn, error) {

	attempt := 1

	for {
		conn, err := clickhouse.Open(&clickhouse.Options{
			Addr: []string{"localhost:9000"},
			Auth: clickhouse.Auth{
				Database: cfg.CLICKHOUSE_DATABASE,
				Username: cfg.CLICKHOUSE_USERNAME,
				Password: cfg.CLICKHOUSE_PASSWORD,
			},
		})

		if err == nil {
			err = conn.Ping(context.Background())
		}

		if err != nil {
			return nil, err
		}

		if err == nil {
			fmt.Print("connected to clickhouse!\n")
			return conn, nil
		}

		attempt++

		if attempt == 5 {
			fmt.Print("Retrying in 5 seconds ...")
			time.Sleep(5 * time.Second)
			attempt = 0
		}

	}

}
