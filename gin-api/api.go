package main

import (
	"go-api/gin-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	api := gin.Default()

	routes.AppRoutes(api)

	// simple endpoint
	// api.GET("/hello", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, "hello-world")
	// })

	api.Run("localhost:3001")
}
