package postgres

import (
	"fmt"
	"log"
	"os"
	"time"

	con "apiGO/run/constLog"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func Connect() (*sqlx.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		log.Fatal("Ошибка загрузки .env файла")
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSLMODE")
	dbPort := os.Getenv("DB_PORT")
	logger, _ := zap.NewDevelopment()
	if err := logger.Sync(); err != nil {
		zap.Error(err)
	}
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s port=%s",
		user, password, dbName, sslMode, dbPort)
	conn, err := sqlx.Connect("postgres", connStr)
	log.Println(connStr)
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
