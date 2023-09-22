package router

import (
	ro "cts/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	group := Router.Group("")
	ro.InitFaceRouter(group)
	return Router
}
