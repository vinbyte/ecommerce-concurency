package postgres

import (
	"context"
	"database/sql"
	"strconv"
)

// CheckUserByID check user by id
func (pr *postgresRepository) CheckUserByID(ctx context.Context, dbTx *sql.Tx, userID int) (int, error) {
	var result int
	query := `select id from users where id = ` + strconv.Itoa(userID)
	var err error
	if dbTx != nil {
		err = dbTx.QueryRowContext(ctx, query).Scan(&result)
	} else {
		err = pr.pgConn.QueryRowContext(ctx, query).Scan(&result)
	}
	if err != nil {
		return result, err
	}
	return result, nil
}
