package postgres

import (
	"database/sql"
	"ecommerce-app/app/helpers"
	"ecommerce-app/domain"
)

type postgresRepository struct {
	pgConn *sql.DB
	helper *helpers.Helpers
}

func NewPostgresRepository(pg *sql.DB, h *helpers.Helpers) domain.OrderRepository {
	return &postgresRepository{
		pgConn: pg,
		helper: h,
	}
}
