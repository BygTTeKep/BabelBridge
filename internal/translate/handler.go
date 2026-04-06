package translate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"babelbridge/internal/shared/errors"
	"babelbridge/internal/translate/domain/entities"
)

type TranslateHandler struct {
	valid *validator.Validate
}

func NewTranslateHandler(valid *validator.Validate) *TranslateHandler {
	return &TranslateHandler{valid: valid}
}

func (th *TranslateHandler) Translate(ctx *gin.Context) {
	var message *entities.MessageEntity
	if err := ctx.ShouldBindJSON(message); err != nil {
		errors.GenerateHTTPError(ctx, errors.NewBadRequest(err.Error()))
	}
	err := th.valid.Struct(message)
	if err != nil {
		errors.GenerateHTTPError(ctx, errors.NewBadRequest(err.Error()))
	}
}
