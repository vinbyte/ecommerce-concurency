package postgres

import (
	"context"
	"database/sql"
)

func (pr *postgresRepository) CleanCartAndOrders(ctx context.Context) error {
	var id int
	query := "delete from cart;delete from orders"
	err := pr.pgConn.QueryRowContext(ctx, query).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}
