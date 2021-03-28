package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

func (pr *postgresRepository) DeleteCart(ctx context.Context, dbTx *sql.Tx, cartID int) error {
	var id int
	query := fmt.Sprintf("delete from cart where id = %d", cartID)
	var err error
	if dbTx != nil {
		err = dbTx.QueryRowContext(ctx, query).Scan(&id)
	} else {
		err = pr.pgConn.QueryRowContext(ctx, query).Scan(&id)
	}
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}
