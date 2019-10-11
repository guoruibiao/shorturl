package main

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/guoruibiao/shorturl/service"
)

func main() {
	app := gin.Default()
	shortURLService, _ := service.New()

	app.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	app.GET("/", func(ctx *gin.Context) {
		result := "service format: /shorturl?longurl=https://github.com"
		ctx.JSON(http.StatusOK, gin.H{"ret": result})
	})

	app.GET("/shorturl", func(ctx *gin.Context) {
		longurl := ctx.DefaultQuery("longurl", "")
		longurl = url.QueryEscape(longurl)
		ret, err := shortURLService.ShortURL(longurl)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"ret": nil})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"ret": ret})
		}
	})

	app.Run(":8080")
}
