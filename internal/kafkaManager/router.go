package kafkamanager

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"babelbridge/internal/kafkamanager/repositories"
	"babelbridge/internal/shared/errors"
	"babelbridge/internal/shared/middlewares"
)

type KafkaManagerHandler struct {
	services IKafkaManagerServices
}

func NewKafkaManagerHandler(services IKafkaManagerServices) *KafkaManagerHandler {
	return &KafkaManagerHandler{services}
}

// CreateTopicHandle method  
func (kmh *KafkaManagerHandler) CreateTopicHandle(ctx *gin.Context) {
	var t repositories.Topic

	if err := ctx.ShouldBindJSON(&t); err != nil {
		errors.GenerateHTTPError(ctx, errors.NewBadRequest(err.Error()))
	}
	token := ctx.GetHeader("X-API-TOKEN")
	if err := kmh.services.CreateTopic(ctx.Request.Context(), t, token); err != nil {
		errors.GenerateHTTPError(ctx, err)
	}
	ctx.Status(http.StatusCreated)
}

// KafkaManagerRouter function  
func KafkaManagerRouter(rg *gin.RouterGroup, handle *KafkaManagerHandler) {
	km := rg.Group("/kafka")
	// km.Use(func(ctx *gin.Context) {
	// 	authHeader := ctx.GetHeader("Authorization")
	// 	fmt.Println(authHeader)
	// })
	km.POST("/topic", middlewares.AuthMiddleware, handle.CreateTopicHandle)
	km.DELETE("/topic/:topic/delete", middlewares.AuthMiddleware, func(ctx *gin.Context) {})
}
