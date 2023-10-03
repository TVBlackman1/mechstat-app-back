package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	app := gin.Default()
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello world")
	})

	return app
}
