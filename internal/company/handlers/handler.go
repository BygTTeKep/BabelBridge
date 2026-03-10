package handlers

import (
	"github.com/gin-gonic/gin"

	"babelbridge/internal/company/services"
)

type Handler struct {
	services services.ICreateCompanyService
}

func NewCompanyHandler(s services.ICreateCompanyService) *Handler {
	return &Handler{
		services: s,
	}
}

func (h *Handler) CreateCompanyHandler(ctx *gin.Context) {
	h.services.Create(ctx)
}
