package internal

import (
	"github.com/gin-gonic/gin"

	"babelbridge/internal/company"
	kafkamanager "babelbridge/internal/kafkaManager"
)

type Routers struct {
	rg      *gin.Engine
	service Services
}

func NewRouters(
	rg *gin.Engine,
	service Services,
) *Routers {
	return &Routers{rg, service}
}

func (r *Routers) Init() {
	api := r.rg.Group("/api")
	comppanyHandler := company.NewCompanyHandler(r.service.CompanyService)
	kafkaManagerHandlers := kafkamanager.NewKafkaManagerHandler(r.service.KafkaManagerService)
	companyRoutes := company.NewCompanyRouter(comppanyHandler)

	kafkamanager.KafkaManagerRouter(api, kafkaManagerHandlers)
	companyRoutes.CompanyRouter(api)
}
