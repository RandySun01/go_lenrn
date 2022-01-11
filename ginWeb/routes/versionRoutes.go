package routes

import (
	"ginWeb/settings"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
@author RandySun
@create 2022-01-11-22:00
*/
func VersionRouters(r *gin.Engine) {

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Conf.Version)
	})
}
