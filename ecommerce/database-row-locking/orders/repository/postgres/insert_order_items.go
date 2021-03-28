package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

// InsertOrderItems is inserting into order_items table
func (pr *postgresRepository) InsertOrderItems(ctx context.Context, dbTx *sql.Tx, orderID int, productCode string, qty int) error {
	var id int
	query := fmt.Sprintf("insert into order_items (product_code,quantity,order_id) values ('%s',%d,%d)", productCode, qty, orderID)
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
