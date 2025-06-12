package main

import (
	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/internal/initialize"

	_ "github.com/anonystick/go-drunk-backend-api-by-ddd-java/cmd/swag/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Drunk Backend API by DDD
// @version 1.0
// @description This is a server for a Go Drunk Backend API, demonstrating DDD principles.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8800
// @BasePath /v1/2025

// @externalDocs.description OpenAPI
// @externalDocs.url https://swagger.io/resources/open-api/
func main() {
	r, port := initialize.Run()

	// prometheus.MustRegister(pingCounter)

	// r.GET("/ping/200", ping)
	// r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":" + port) // listen and serve on :8899
}
