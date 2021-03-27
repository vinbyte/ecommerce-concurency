package helpers

import (
	"database/sql"
	"time"
)

// Helpers ...
type Helpers struct {
	logMaxAge time.Duration
	dbConn    *sql.DB
	dbTrx     *sql.Tx
}

// New making new instance of this library
func New() *Helpers {
	lpc := &Helpers{}
	return lpc
}

func (h *Helpers) SetLogMaxAge(age time.Duration) {
	h.logMaxAge = age
}
