package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

func (pr *postgresRepository) InsertCartItems(ctx context.Context, cartID int, productCode string, qty int) error {
	var id int
	query := fmt.Sprintf(`insert into cart_items (cart_id, product_code, quantity) values (%d, '%s', %d)`, cartID, productCode, qty)
	err := pr.helper.QueryRowContext(ctx, query).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}
