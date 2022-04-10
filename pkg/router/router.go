package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rohandas-max/urlShortner/pkg/controller"
)

func Router() *gin.Engine {
	router := gin.Default()
	err := router.SetTrustedProxies([]string{"0.0.0.0"})
	if err != nil {
		panic(err)
	}
	router.POST("/shorturl", controller.UrlShortener)
	router.GET("/", controller.VisitShortURL)
	router.NoRoute(controller.Default)
	return router
}
