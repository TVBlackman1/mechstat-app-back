package handlers

import (
	"encoding/json"
	"mechstat/internal/demo"
	"net/http"
	"strconv"

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
		rawPageNumber := ctx.DefaultQuery("page", "1")
		pageNumber, err := strconv.Atoi(rawPageNumber)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
		}
		// TODO change to default gin using
		res, _ := json.Marshal(demo.GetExperimentsPagination(pageNumber))
		ctx.Writer.Header().Add("Content-Type", "application/json")
		ctx.Writer.Write(res)
	})

	return app
}
