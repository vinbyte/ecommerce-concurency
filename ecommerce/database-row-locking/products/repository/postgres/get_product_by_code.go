package postgres

import (
	"context"
	"ecommerce-app/domain"
)

// GetProductByCode get the product using product code
func (pr *postgresRepository) GetProductByCode(ctx context.Context, productCode string, isRowLocking bool) (domain.Product, error) {
	var result domain.Product
	query := `select id, code, name, description, stock, price from products where code = '` + productCode + `'`
	if isRowLocking {
		query += " for update"
	}
	err := pr.helper.QueryRowContext(ctx, query).Scan(&result.ID, &result.Code, &result.Name, &result.Desc, &result.Stock, &result.Price)
	if err != nil {
		return result, err
	}
	return result, nil
}
