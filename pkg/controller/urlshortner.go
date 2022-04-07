package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UrlShortner(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello",
	})
}
