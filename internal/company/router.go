package company

import (
	"github.com/gin-gonic/gin"

	"babelbridge/internal/company/handlers"
)

type CompanyRoute struct {
	handler handlers.Handler
}

func (cr *CompanyRoute) CreateCompany(ctx *gin.Context) {
	cr.services.Create(ctx)
}

func NewCompanyRoutes(services ICompanyServices) *CompanyRoute {
	return &CompanyRoute{services}
}

func CompanyRouter(rg *gin.RouterGroup, handle CompanyRoute) {
	company := rg.Group("/company")
	company.POST("/new", handle.CreateCompany)
}
