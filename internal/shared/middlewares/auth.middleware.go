package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	header := ctx.GetHeader("X-API-TOKEN")
	if len(header) == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": ""})
	}
}
