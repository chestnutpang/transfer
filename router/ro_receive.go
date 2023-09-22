package router

import (
	"cts/apps/receive"

	"github.com/gin-gonic/gin"
)

func InitFaceRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("/cts")

	{
		baseRouter.POST("/receive", receive.Receive)
	}
}
