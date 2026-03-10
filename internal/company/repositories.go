package company

import (
	"database/sql"

	"babelbridge/internal/company/repositories"
)

type ICompanyRepository interface {
	repositories.ICreateCompanyRepository
}

type CompanyRepository struct {
	repositories.ICreateCompanyRepository
}

func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{
		ICreateCompanyRepository: repositories.NewCreateCompanyRepository(db),
	}
}
