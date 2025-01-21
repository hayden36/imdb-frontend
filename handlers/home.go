package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hayden36/imdb-frontend/views"
)

func HandleHomeRoute() func(c *gin.Context) {
	return func(c *gin.Context) {
		views.HomeView().Render(c, c.Writer)
	}
}
