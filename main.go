package main

import (
	"weather/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.POST("/weather", handler.Weather)

	router.Run(":8080")
}
