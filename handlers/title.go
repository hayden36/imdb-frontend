package handlers

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/hayden36/imdb-frontend/structs"
	"github.com/hayden36/imdb-frontend/views"
	"log"
	"net/http"
)

func HandleTitleRoute() func(c *gin.Context) {
	return func(c *gin.Context) {
		titleParameter := c.Param("title")
		req, _ := http.NewRequest("GET", "https://imdb.com/title/"+titleParameter, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:133.0) Gecko/20100101 Firefox/133.0")

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			log.Fatal(err)
		}
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		nameSelection := doc.Find("span.hero__primary-text").First()
		plotSelection := doc.Find("p[data-testid=plot]").Find("span").First()
		var topCastSlice []structs.Cast
		doc.Find("div[data-testid=title-cast-item]").Each(func(i int, s *goquery.Selection) {
			castNameSelection := s.Find("a")
			person := structs.Cast{Name: castNameSelection.AttrOr("aria-label", "no name"), Link: castNameSelection.AttrOr("href", "#")}
			topCastSlice = append(topCastSlice, person)
		})

		movie := structs.MovieShow{Name: nameSelection.Text(), Plot: plotSelection.Text(), TopCast: topCastSlice}

		views.TitleView(movie).Render(c, c.Writer)

	}
}
