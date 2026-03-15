package company

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

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
	query := `
		INSERT INTO company(name, token)
		VALUES ($1, $2)
		RETURNING id
	`
	for range 3 {
		token, err := uuid.NewV7()
		if err != nil {
			return err
		}
		var id int
		err = c.db.QueryRowContext(ctx, query, dto.Name, token).Scan(&id)
		if err == nil {
			return nil
		}
	}
	return fmt.Errorf("failed to generate unique token")
}

// Delete method  
func (c *CompanyRepository) Delete(ctx context.Context, id, token string) error {
	query := `DELETE FROM company WHERE id = $1 AND token = $2`

	res, err := c.db.ExecContext(ctx, query, id, token)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("company not found")
	}

	return nil
}
