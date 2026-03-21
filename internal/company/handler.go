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
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	data, err := h.services.Create(ctx, dto)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"error": "Internal error"})
	}
	ctx.JSON(201, data)
}

func (h *Handler) DeleteComnyByID(ctx *gin.Context) {
	id := ctx.Param("id")
	token := ctx.GetHeader("X-API-TOKEN")
	err := h.services.Delete(ctx.Request.Context(), id, token)
	if err != nil {
		switch err.(type) {
		case *errors.NotFoundError:
			ctx.JSON(404, gin.H{"error": err.Error()})
		case *errors.BadRequestError:
			ctx.JSON(400, gin.H{"error": err.Error()})
		default:
			ctx.JSON(500, gin.H{"error": "internal server error"})
		}
		return
	}
	ctx.Status(http.StatusOK)
}
