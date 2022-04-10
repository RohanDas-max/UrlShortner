package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohandas-max/urlShortner/pkg/database"
)

func VisitShortURL(c *gin.Context) {
	url := c.Query("url")
	v := database.GetUrl(url)
	c.Redirect(http.StatusMovedPermanently, v)
}
