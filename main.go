package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hayden36/imdb-frontend/handlers"
	"github.com/hayden36/imdb-frontend/pages"
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
	router.GET("/search", handlers.SearchHandler())

	router.NoRoute(func(c *gin.Context) {
		pages.Error(404, "Page not found.").Render(c, c.Writer)
	})

	router.Run("127.0.0.1:3000")
}
