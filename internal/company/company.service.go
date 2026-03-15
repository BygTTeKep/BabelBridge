package company

import (
	"context"

	"babelbridge/internal/company/dtos"
)

type ICompanyService interface {
	Create(ctx context.Context, dto *dtos.CreateCompanyDto) error
	Delete(ctx context.Context, id, token string) error
}

type CompanyService struct {
	repo ICompanyRepository
}

func NewCompanyService(repo ICompanyRepository) *CompanyService {
	return &CompanyService{
		repo: repo,
	}
}

func (c *CompanyService) Create(ctx context.Context, dto *dtos.CreateCompanyDto) error {
	if err := c.repo.Create(ctx, dto); err != nil {
		return err
	}
	return nil
}

func (c *CompanyService) Delete(ctx context.Context, id, token string) error {
	if err := c.repo.Delete(ctx, id, token); err != nil {
		return err
	}
	return nil
}
