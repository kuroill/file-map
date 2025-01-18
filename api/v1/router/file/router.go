package file

import (
	"file-map-server/api/v1/handler/file"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
	*file.Handler
}

func New(e *gin.Engine) *Router {
	r := &Router{e, file.New()}
	r.SetupFile()
	return r
}
