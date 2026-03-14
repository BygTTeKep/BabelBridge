package company

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"babelbridge/internal/company/dtos"
)

type Handler struct {
	services ICompanyService
}

func NewCompanyHandler(s ICompanyService) *Handler {
	return &Handler{
		services: s,
	}
}

func (h *Handler) CreateCompanyHandler(ctx *gin.Context) {
	dto := &dtos.CreateCompanyDto{}
	if err := ctx.ShouldBind(dto); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	if err := h.services.Create(ctx, dto); err != nil {
		ctx.JSON(500, gin.H{"error": "Internal error"})
	}
	ctx.JSON(201, dto) // или другой ответ
}

func (h *Handler) DeleteComnyByID(ctx *gin.Context) {
	id := ctx.Param("id")
	token := ctx.GetHeader("X-API-TOKEN")
	if err := h.services.Delete(ctx.Request.Context(), id, token); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
	}
	ctx.Status(http.StatusOK)
}
