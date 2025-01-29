package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hayden36/imdb-frontend/pages"
)

func HandleHomeRoute() func(c *gin.Context) {
	return func(c *gin.Context) {
		pages.HomeView().Render(c, c.Writer)
	}
}
