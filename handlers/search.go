package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hayden36/imdb-frontend/pages"
	"github.com/hayden36/imdb-frontend/structs"
	"github.com/hayden36/imdb-frontend/utils"
)

func SearchHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		searchQuery := c.Query("q")
		// names: https://v3.sg.media-imdb.com/suggestion/names/x/jimmy.json?includeVideos=1
		req, _ := http.NewRequest("GET", "https://v3.sg.media-imdb.com/suggestion/x/"+searchQuery+".json?includeVideos=0", nil)
		req.Header.Set("User-Agent", utils.USER_AGENT)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			pages.Error(500, "Unable to send request.").Render(c, c.Writer)
			log.Print(err)
			return
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			pages.Error(res.StatusCode, fmt.Sprintf("Error from IMDB: %s", res.Status)).Render(c, c.Writer)
			log.Printf("status code error: %d %s", res.StatusCode, res.Status)
			return
		}
		response := structs.SearchResponse{}
		body, _ := io.ReadAll(res.Body)
		err = json.Unmarshal(body, &response)
		if err != nil {
			pages.Error(500, "Error decoding JSON.").Render(c, c.Writer)
			return
		}

		pages.SearchPage(response.Query, response).Render(c, c.Writer)
	}
}
