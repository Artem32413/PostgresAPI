package postgres

import (
	"fmt"
	"time"

	con "apiGO/run/constLog"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func Connect() (*sqlx.DB, error) {
	logger, _ := zap.NewDevelopment()
	if err := logger.Sync(); err != nil {
		zap.Error(err)
	}
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	conn, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		logger.Error(con.ErrDBConnect,
			zap.Error(err),
			zap.Duration("backoff", time.Second),
		)
		return nil, fmt.Errorf("sqlx connect")
	}
	err = conn.Ping()
	if err != nil {
		logger.Error(con.ErrDBPing,
			zap.Error(err),
			zap.Duration("backoff", time.Second),
		)
		return nil, fmt.Errorf("ping error")
	}
	return conn, nil
}
