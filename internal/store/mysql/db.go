package mysql

import (
	"database/sql"
	"fmt"
	"os"
	// _ "github.com/go-sql-driver/mysql"
)

// OpenFromEnv opens *sql.DB using MYSQL_DSN env variable.
// Example: app:app@tcp(db:3306)/projcardapio?parseTime=true&charset=utf8mb4&loc=Local
func OpenFromEnv() (*sql.DB, error) {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "app:app@tcp(db:3306)/projcardapio?parseTime=true&charset=utf8mb4&loc=Local"
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql open: %w", err)
	}
	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping: %w", err)
	}
	return db, nil
}
