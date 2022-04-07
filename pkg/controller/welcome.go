package controller

import (
	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	c.Redirect(200, "localhost:8080/url-shortner")
}
