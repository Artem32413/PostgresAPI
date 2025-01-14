package postgres

import (
	"fmt"
	"os"
	"time"

	con "apiGO/run/constLog"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func Connect() (*sqlx.DB, error) {
	logger, _ := zap.NewDevelopment()
	if err := logger.Sync(); err != nil {
		zap.Error(err)
	}
	err := godotenv.Load()
	if err != nil {
		logger.Error(con.ErrEnv,
			zap.Error(err))
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSLMODE")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s port=%s host=%s",
		user, password, dbName, sslMode, dbPort, dbHost)
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
			zap.Error(err))
		return nil, fmt.Errorf("ping error")
	}
	return conn, nil
}
