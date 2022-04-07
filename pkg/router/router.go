package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rohandas-max/urlShortner/pkg/controller"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"0.0.0.0"})

	r.GET("/", controller.Welcome)
	r.POST("/url-shortner", controller.UrlShortner)
	r.NoRoute(controller.Default)
	return r
}
