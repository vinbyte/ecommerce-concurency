package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

func (pr *postgresRepository) DeleteCart(ctx context.Context, cartID int) error {
	var id int
	query := fmt.Sprintf("delete from cart where id = %d", cartID)
	err := pr.helper.QueryRowContext(ctx, query).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}
