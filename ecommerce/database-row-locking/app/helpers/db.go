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
	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(25)
	conn.SetConnMaxLifetime(5 * time.Minute)
	h.dbConn = conn
	return conn
}

func (h *Helpers) BeginTrx(ctx context.Context, opts *sql.TxOptions) error {
	trx, err := h.dbConn.BeginTx(ctx, opts)
	if err == nil {
		h.dbTrx = trx
	}
	return err
}

func (h *Helpers) CommitTrx() error {
	return h.dbTrx.Commit()
}

func (h *Helpers) RollbackTrx() error {
	return h.dbTrx.Rollback()
}

func (h *Helpers) QueryContext(ctx context.Context, query string) (rows *sql.Rows, err error) {
	if h.dbTrx != nil {
		rows, err = h.dbTrx.QueryContext(ctx, query)
	} else {
		rows, err = h.dbConn.QueryContext(ctx, query)
	}
	return
}

func (h *Helpers) QueryRowContext(ctx context.Context, query string) (row *sql.Row) {
	if h.dbTrx != nil {
		row = h.dbTrx.QueryRowContext(ctx, query)
	} else {
		row = h.dbConn.QueryRowContext(ctx, query)
	}
	return
}
