package server

import (
	"cts/init/log"
	"cts/init/router"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	log.Init()

	rootRouter := router.Routers()
	//address := ":9999"
	rootRouter.GET("/", func(c *gin.Context) {
		c.String(200, "Hello world.")
	})
	//rootRouter.Run(address)
	return rootRouter
}
