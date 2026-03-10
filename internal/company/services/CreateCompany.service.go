package services

import (
	"context"

	"github.com/gin-gonic/gin"

	"babelbridge/internal/company/dtos"
	"babelbridge/internal/company/repositories"
)

type ICreateCompanyService interface {
	Create(ctx context.Context)
}

type CreateCompanyService struct {
	repo repositories.ICreateCompanyRepository
}

func NewCreateCompanyService(repo repositories.ICreateCompanyRepository) *CreateCompanyService {
	return &CreateCompanyService{
		repo: repo,
	}
}

func (c *CreateCompanyService) Create(ctx context.Context) {
	dto := &dtos.CreateCompanyDto{}
	if err := ctx.ShouldBind(dto); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	if err := c.repo.Create(ctx, dto); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, dto) // или другой ответ
}
