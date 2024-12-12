package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() (*sqlx.DB, error) {
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	conn, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("sqlx connect")
	}
	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping error")
	}
	return conn, nil
}
