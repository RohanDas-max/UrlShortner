package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohandas-max/urlShortner/pkg/database"
	"github.com/rohandas-max/urlShortner/pkg/handler"
	"github.com/rohandas-max/urlShortner/pkg/platform"
)

func UrlShortener(c *gin.Context) {
	var r platform.UrlReqBody
	if err := c.ShouldBindJSON(&r); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Binding error " + err.Error(),
		})
		return
	}
	if exist := database.IsExist(r.Url); exist {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorMessage": "url already existed",
		})
		return
	}

	rand := handler.RandomURL(4)
	database.InsertUrl(r, rand)
	c.JSON(http.StatusOK, gin.H{
		"shortUrl": "localhost:8080?url=" + rand,
	})

}
