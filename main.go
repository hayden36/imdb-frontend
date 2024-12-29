package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hayden36/imdb-frontend/handlers"
)

func main() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.GET("/name/:name", handlers.HandleNameRoute())
	router.Run(":3000")
}
