package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hayden36/imdb-frontend/handlers"
)

// https://xeiaso.net/blog/using-tailwind-go/
//
//go:generate npm run build
//go:generate templ generate
func main() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.GET("/", handlers.HandleHomeRoute())
	router.GET("/name/:name", handlers.HandleNameRoute())
	router.GET("/title/:title", handlers.HandleTitleRoute())
	router.Run("0.0.0.0:3000")
}
