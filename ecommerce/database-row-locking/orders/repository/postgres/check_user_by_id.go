package postgres

import (
	"context"
	"strconv"
)

// CheckUserByID check user by id
func (pr *postgresRepository) CheckUserByID(ctx context.Context, userID int) (int, error) {
	var result int
	query := `select id from users where id = ` + strconv.Itoa(userID)
	err := pr.helper.QueryRowContext(ctx, query).Scan(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
