package helpers

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

// InitPostgres is used to initiate to db postgres connection
func (h *Helpers) InitPostgres() *sql.DB {
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	sslMode := os.Getenv("POSTGRES_SSL_MODE")
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, sslMode)
	conn, err := sql.Open(`postgres`, connStr)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	conn.SetMaxOpenConns(50)
	conn.SetMaxIdleConns(50)
	conn.SetConnMaxLifetime(5 * time.Minute)
	h.dbConn = conn
	return conn
}

// BeginTrx is helper to begin the database transaction
func (h *Helpers) BeginTrx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	trx, err := h.dbConn.BeginTx(ctx, opts)
	return trx, err
}
