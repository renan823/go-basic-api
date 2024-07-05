package main

import (
	"github.com/gin-gonic/gin"

	controllers "files/controllers"
)

func main () {
	app := gin.Default()

	app.GET("/", controllers.FindAll)
	app.GET("/:name", controllers.FindByName)
	app.POST("/upload", controllers.Upload)

	app.Run(":3000")
}