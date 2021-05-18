package main

import (
	"github.com/gin-gonic/gin"
	"github.com/josesalasdev/golang_api_template/src/controllers"
	"github.com/josesalasdev/golang_api_template/src/models"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	// routes
	r.GET("/ping", controllers.Ping)
	r.GET("/books", controllers.FindBooks)

	r.Run() // listen and serve on 0.0.0.0:8080
}
