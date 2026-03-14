package company

import (
	"context"
	"database/sql"
	"fmt"

	"babelbridge/internal/company/dtos"
)

type ICompanyRepository interface {
	Create(ctx context.Context, dto *dtos.CreateCompanyDto) error
	Delete(ctx context.Context, id, token string) error
}

type CompanyRepository struct {
	db *sql.DB
}

func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{
		db: db,
	}
}

// Create method  
func (c *CompanyRepository) Create(ctx context.Context, dto *dtos.CreateCompanyDto) error {
	var id int
	// TODO
	// генерим токен
	// token := ""
	query := fmt.Sprintf("INSERT INTO %s(name) VALUES(%s) RETURNING id", "company", dto.Name)
	row := c.db.QueryRowContext(ctx, query)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

// Delete method  
func (c *CompanyRepository) Delete(ctx context.Context, id, token string) error {
	return nil
}
