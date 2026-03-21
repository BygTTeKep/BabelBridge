package company

// TODO добавить нормальные логи

import (
	"context"

	"babelbridge/internal/company/dtos"

	log "github.com/sirupsen/logrus"
)

type ICompanyService interface {
	Create(ctx context.Context, dto *dtos.CreateCompanyDto) (*dtos.CreateCompanyResponseDto, error)
	Delete(ctx context.Context, id, token string) error
}

type CompanyService struct {
	repo   ICompanyRepository
	logger *log.Entry
}

func NewCompanyService(repo ICompanyRepository, logger *log.Logger) *CompanyService {
	return &CompanyService{
		repo:   repo,
		logger: logger.WithField("service", "CompanyService"),
	}
}

func (c *CompanyService) Create(ctx context.Context, dto *dtos.CreateCompanyDto) (*dtos.CreateCompanyResponseDto, error) {
	data, err := c.repo.Create(ctx, dto)
	if err != nil {
		c.logger.Errorf("failed to create cimpany: %v", err)
		return nil, err
	}
	return data, nil
}

func (c *CompanyService) Delete(ctx context.Context, id, token string) error {
	if err := c.repo.Delete(ctx, id, token); err != nil {
		c.logger.Errorf("failed to delete comapny: %v", err)
		return err
	}
	return nil
}
