package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/larksuite/oapi-sdk-go/core/config"
	"github.com/larksuite/oapi-sdk-go/event/http"
)

func Register(path string, conf *config.Config, g *gin.Engine) {
	g.POST(path, func(context *gin.Context) {
		http.Handle(conf, context.Request, context.Writer)
	})
}
