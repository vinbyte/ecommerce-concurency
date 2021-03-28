package postgres

import (
	"context"
	"strconv"
)

func (pr *postgresRepository) CreateCart(ctx context.Context, userID int) (int, error) {
	var result int
	query := `insert into cart (user_id) values (` + strconv.Itoa(userID) + `) RETURNING id`
	err := pr.helper.QueryRowContext(ctx, query).Scan(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
