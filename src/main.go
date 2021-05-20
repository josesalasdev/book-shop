package main

import (
	"github.com/gin-gonic/gin"
	"github.com/josesalasdev/golang_api_template/src/controllers"
	"github.com/josesalasdev/golang_api_template/src/models"
	"github.com/josesalasdev/golang_api_template/docs"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/files" // swagger embed files
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	// routes
	r.GET("/v1/ping", controllers.Ping)
	r.GET("/v1/books", controllers.FindBooks)

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server golang."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run() // listen and serve on 0.0.0.0:8080
}
