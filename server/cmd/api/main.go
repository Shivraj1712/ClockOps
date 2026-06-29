package main

import (
	"net/http"

	_ "github.com/Shivraj1712/ClockOps.git/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// HealthCheck godoc
// @Summary         First Route
// @Description     testing if the server is running or not
// @Tags            Health
// @Accept          json
// @Produce         json
// @Success         200 {object} map[string]string
// @Failure         500 {object} map[string]string
// @Router          / [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Server Running on Port : 8080",
	})
}

// @title			ClockOps API
// @version 		1.0
// @description 	HR management APi
// @host 			localhost:8080
// @BasePath 		/
func main() {
	r := gin.Default()

	r.GET("/", HealthCheck)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
