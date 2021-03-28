package postgres

import (
	"context"
	"ecommerce-app/domain"
	"strconv"
)

// GetAllProducts is query for getting product list
func (pr *postgresRepository) GetAllProducts(ctx context.Context, offset int, limit int) ([]domain.Product, error) {
	result := make([]domain.Product, 0)
	query := `select id, code, name, description, stock, price from products 
	order by name asc 
	limit ` + strconv.Itoa(limit) + ` 
	offset ` + strconv.Itoa(offset)
	// I use pr.helper instead of pr.pgConn because we need to know, there is any transaction happen or not. Also, we can easily rollback if any error happen in usecase.
	rows, err := pr.helper.QueryContext(ctx, query)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		temp := domain.Product{}
		err := rows.Scan(&temp.ID, &temp.Code, &temp.Name, &temp.Desc, &temp.Stock, &temp.Price)
		if err != nil {
			return result, err
		}
		result = append(result, temp)
	}
	return result, nil
}
