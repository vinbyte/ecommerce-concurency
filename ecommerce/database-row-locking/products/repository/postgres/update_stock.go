package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

// UpdateStock is for updating stock at checkout process
func (pr *postgresRepository) UpdateStock(ctx context.Context, dbTx *sql.Tx, productCode string, stock int) error {
	var id int
	query := fmt.Sprintf("update products set stock = %d where code = '%s'", stock, productCode)
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
