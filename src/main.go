package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/josesalasdev/book-shop/docs"
	"github.com/josesalasdev/book-shop/src/config"
	"github.com/josesalasdev/book-shop/src/controllers"
	"github.com/josesalasdev/book-shop/src/middlewares"
	"github.com/josesalasdev/book-shop/src/models"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT
func main() {
	r := gin.Default()

	models.ConnectDataBase()

	// routes
	r.GET("/v1/ping", controllers.Ping)
	r.POST("v1/signin", controllers.SignIn)

	apiUser := r.Group("/v1/user")
	{
		apiUser.POST("/", controllers.CreateUser)
	}

	apiBook := r.Group("/v1/book")
	apiBook.Use(middlewares.AuthorizeJWT())
	{
		apiBook.POST("/", controllers.CreateBook)
		apiBook.GET("/", controllers.FindBooks)
		apiBook.GET("/:id", controllers.DetailBooks)
	}

	apiCategory := r.Group("/v1/category")
	{
		apiCategory.POST("/", controllers.CreateCategory)
		apiCategory.GET("/", controllers.FindCategory)
	}

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Book Shop API"
	docs.SwaggerInfo.Description = "This is a sample server golang."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(config.Port) // nolint
}
