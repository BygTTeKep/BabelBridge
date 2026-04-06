package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"babelbridge/internal/company"
	kafkamanager "babelbridge/internal/kafkaManager"
	"babelbridge/internal/translate"
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

var validate *validator.Validate

func (r *Routers) Init() {
	api := r.rg.Group("/api")
	comppanyHandler := company.NewCompanyHandler(r.service.CompanyService)
	kafkaManagerHandlers := kafkamanager.NewKafkaManagerHandler(r.service.KafkaManagerService)
	translateHandlers := translate.NewTranslateHandler(validate)

	translateRouter := translate.NewTranslateRouter(translateHandlers)
	companyRoutes := company.NewCompanyRouter(comppanyHandler)
	kafkamanager.KafkaManagerRouter(api, kafkaManagerHandlers)
	companyRoutes.CompanyRouter(api)
	translateRouter.TranslateRouter(api)
}
