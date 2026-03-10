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
	kafkaManagerHandlers := kafkamanager.NewKafkaManagerHandler(r.service.KafkaManagerService)
	companyRoutes := company.NewCompanyRoutes(r.service.CompanyService)

	kafkamanager.KafkaManagerRouter(api, kafkaManagerHandlers)
	company.CompanyRouter(api, *companyRoutes)
}
