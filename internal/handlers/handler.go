package handlers

import (
	"encoding/json"
	"mechstat/internal/demo"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	app := gin.Default()
	app.Use(cors.Default())
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello world")
	})

	app.GET("/experiments", func(ctx *gin.Context) {
		// TODO change to default gin using
		res, _ := json.Marshal(demo.GetExperimentsRandomly(30))
		ctx.Writer.Header().Add("Content-Type", "application/json")
		ctx.Writer.Write(res)
	})

	return app
}
