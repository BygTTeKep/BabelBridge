package company

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"babelbridge/internal/company/dtos"
	"babelbridge/internal/shared/errors"
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
		log.Println(err)
		errors.GenerateHTTPError(ctx, errors.NewBadRequest(err.Error()))
	}
	data, err := h.services.Create(ctx, dto)
	if err != nil {
		log.Println(err)
		errors.GenerateHTTPError(ctx, err)
	}
	ctx.JSON(http.StatusCreated, data)
}

func (h *Handler) DeleteComnyByID(ctx *gin.Context) {
	id := ctx.Param("id")
	token := ctx.GetHeader("X-API-TOKEN")
	err := h.services.Delete(ctx.Request.Context(), id, token)
	if err != nil {
		errors.GenerateHTTPError(ctx, err)
	}
	ctx.Status(http.StatusOK)
}
