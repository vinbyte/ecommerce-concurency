package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

// CreateOrders is inserting into orders table
func (pr *postgresRepository) CreateOrders(ctx context.Context, dbTx *sql.Tx, userID int) (int, error) {
	var orderID int
	query := fmt.Sprintf("insert into orders (user_id) values (%d) RETURNING id", userID)
	var err error
	if dbTx != nil {
		err = dbTx.QueryRowContext(ctx, query).Scan(&orderID)
	} else {
		err = pr.pgConn.QueryRowContext(ctx, query).Scan(&orderID)
	}
	if err != nil {
		return orderID, err
	}
	return orderID, nil
}
