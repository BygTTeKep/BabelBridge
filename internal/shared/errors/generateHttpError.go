package errors

import "github.com/gin-gonic/gin"

func GenerateHTTPError(ctx *gin.Context, err error) {
	switch err.(type) {
	case *NotFoundError:
		ctx.JSON(404, gin.H{"error": err.Error()})
	case *BadRequestError:
		ctx.JSON(400, gin.H{"error": err.Error()})
	default:
		ctx.JSON(500, gin.H{"error": "internal server error"})
	}
}
