package postgres

import (
	"context"
	"database/sql"
	"strconv"
)

func (pr *postgresRepository) CreateCart(ctx context.Context, dbTx *sql.Tx, userID int) (int, error) {
	var result int
	query := `insert into cart (user_id) values (` + strconv.Itoa(userID) + `) RETURNING id`
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
