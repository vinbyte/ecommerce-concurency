package postgres

import (
	"context"
	"database/sql"
	"ecommerce-app/domain"
	"fmt"
)

// GetCartData is getting shopping cart data with the details
func (pr *postgresRepository) GetCartData(ctx context.Context, dbTx *sql.Tx, cartID int, isRowLocking bool) (domain.CartData, error) {
	var result domain.CartData
	query := fmt.Sprintf("select id, user_id, date from cart where id = %d", cartID)
	if isRowLocking {
		query += " for update"
	}
	var err error
	if dbTx != nil {
		err = dbTx.QueryRowContext(ctx, query).Scan(&result.CartID, &result.UserID, &result.Date)
	} else {
		err = pr.pgConn.QueryRowContext(ctx, query).Scan(&result.CartID, &result.UserID, &result.Date)
	}
	if err != nil {
		return result, err
	}
	query = fmt.Sprintf("select product_code, quantity, date from cart_items where cart_id = %d", result.CartID)
	if isRowLocking {
		query += " for update"
	}
	var rows *sql.Rows
	if dbTx != nil {
		rows, err = dbTx.QueryContext(ctx, query)
	} else {
		rows, err = pr.pgConn.QueryContext(ctx, query)
	}
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var item domain.CartItem
		err := rows.Scan(&item.ProductCode, &item.Qty, &item.Date)
		if err != nil {
			return result, err
		}
		result.Items = append(result.Items, item)
	}
	return result, nil
}
