package postgres

import (
	"context"
	"ecommerce-app/domain"
	"strconv"
)

func (pr *postgresRepository) GetAllProducts(ctx context.Context, offset int, limit int) ([]domain.Product, error) {
	result := make([]domain.Product, 0)
	query := `select id, code, name, description, stock, price from products 
	order by name asc 
	limit ` + strconv.Itoa(limit) + ` 
	offset ` + strconv.Itoa(offset)
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
