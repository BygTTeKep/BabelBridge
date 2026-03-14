package company

import (
	"github.com/gin-gonic/gin"

	"babelbridge/internal/shared/middlewares"
)

type CompanyRoute struct {
	handler *Handler
}

func NewCompanyRouter(handler *Handler) *CompanyRoute {
	return &CompanyRoute{handler}
}

func (cr *CompanyRoute) CompanyRouter(rg *gin.RouterGroup) {
	company := rg.Group("/company")
	company.POST("/new", middlewares.AuthMiddleware, cr.handler.CreateCompanyHandler)
	company.DELETE(":id/delete", middlewares.AuthMiddleware, cr.handler.DeleteComnyByID)
}
