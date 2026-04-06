package translate

import "github.com/gin-gonic/gin"

type TranslateRoute struct {
	handler *TranslateHandler
}

func NewTranslateRouter(handler *TranslateHandler) *TranslateRoute {
	return &TranslateRoute{
		handler: handler,
	}
}

func (tr *TranslateRoute) TranslateRouter(rg *gin.RouterGroup) {
	rg.POST("", tr.handler.Translate)
}
