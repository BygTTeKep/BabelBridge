package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"babelbridge/internal/company/dtos"
)

type ICreateCompanyRepository interface {
	Create(ctx context.Context, dto *dtos.CreateCompanyDto) error
	Delete(ctx context.Context, id, token string) error
}

type CreateCompanyRepository struct {
	db *sql.DB
}

func NewCreateCompanyRepository(db *sql.DB) *CreateCompanyRepository {
	return &CreateCompanyRepository{
		db: db,
	}
}

// Create method  
func (c *CreateCompanyRepository) Create(ctx context.Context, dto *dtos.CreateCompanyDto) error {
	var id int
	query := fmt.Sprintf("INSERT INTO %s(name) VALUES(%s) RETURNING id", "company", dto.Name)
	row := c.db.QueryRowContext(ctx, query)
	if err := row.Scan(&id); err != nil {
		return err
	}
	fmt.Println(id)
	return nil
}

// Delete method  
func (c *CreateCompanyRepository) Delete(ctx context.Context, id, token string) error {
	return nil
}
