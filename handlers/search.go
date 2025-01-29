package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hayden36/imdb-frontend/pages"
	"github.com/hayden36/imdb-frontend/structs"
	"io"
	"log"
	"net/http"
)

func SearchHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		searchQuery := c.Query("q")
		req, _ := http.NewRequest("GET", "https://v3.sg.media-imdb.com/suggestion/x/"+searchQuery+".json?includeVideos=0", nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:133.0) Gecko/20100101 Firefox/133.0")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}
		response := structs.SearchResponse{}
		body, _ := io.ReadAll(res.Body)
		err = json.Unmarshal(body, &response)

		pages.SearchPage(response.Q, response).Render(c, c.Writer)
	}
}
