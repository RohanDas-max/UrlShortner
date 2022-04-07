package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Default(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "OOPS!! NOT FOUND",
	})
}
