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

// NewPostgresRepository will create an object that represent domain.CourseRepository
func NewPostgresRepository(pg *sql.DB, h *helpers.Helpers) domain.ProductRepository {
	return &postgresRepository{
		pgConn: pg,
		helper: h,
	}
}
