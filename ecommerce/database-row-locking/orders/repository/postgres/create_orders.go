package postgres

import (
	"context"
	"fmt"
)

// CreateOrders is inserting into orders table
func (pr *postgresRepository) CreateOrders(ctx context.Context, userID int) (int, error) {
	var orderID int
	query := fmt.Sprintf("insert into orders (user_id) values (%d) RETURNING id", userID)
	err := pr.helper.QueryRowContext(ctx, query).Scan(&orderID)
	if err != nil {
		return orderID, err
	}
	return orderID, nil
}
