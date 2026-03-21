package company

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"

	"babelbridge/internal/company/dtos"
	"babelbridge/internal/shared/errors"
)

type ICompanyRepository interface {
	Create(ctx context.Context, dto *dtos.CreateCompanyDto) (*dtos.CreateCompanyResponseDto, error)
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
func (c *CompanyRepository) Create(ctx context.Context, dto *dtos.CreateCompanyDto) (*dtos.CreateCompanyResponseDto, error) {
	retryCount := 3
	query := `
		INSERT INTO company(name, token)
		VALUES ($1, $2)
		RETURNING id
	`
	for range retryCount {
		// TODO: шифровать токен
		token, err := uuid.NewV7()
		if err != nil {
			return nil, err
		}
		var id int
		err = c.db.QueryRowContext(ctx, query, dto.Name, token).Scan(&id)
		if err == nil {
			return &dtos.CreateCompanyResponseDto{
					ID:    id,
					Name:  dto.Name,
					Token: token,
				},
				nil
		}
	}
	return nil, fmt.Errorf("failed to generate unique token")
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
		return errors.NewNotFound("company not found")
	}

	return nil
}
