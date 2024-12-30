package handlers

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/hayden36/imdb-frontend/structs"
	"github.com/hayden36/imdb-frontend/views"
	"log"
	"net/http"
)

func HandleNameRoute() func(c *gin.Context) {
	return func(c *gin.Context) {
		nameParameter := c.Param("name")
		req, _ := http.NewRequest("GET", "https://www.imdb.com/name/"+nameParameter, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:133.0) Gecko/20100101 Firefox/133.0")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		nameSelection := doc.Find(".hero__primary-text").First()
		bioSelection := doc.Find(".ipc-html-content-inner-div").First()
		var knownForSlice []structs.KnownFor
		bioHtml, _ := bioSelection.Html()

		doc.Find("div.ipc-metadata-list-summary-item__tc").Each(func(i int, s *goquery.Selection) {
			title := s.Find("a.ipc-metadata-list-summary-item__t").Text()
			link := s.Find("a").AttrOr("href", "none")
			tempKnownFor := structs.KnownFor{Name: title, Link: link}
			knownForSlice = append(knownForSlice, tempKnownFor)
		})

		person := structs.Person{Name: nameSelection.Text(), Bio: bioHtml, KnownFor: knownForSlice}

		views.NameView(person).Render(c, c.Writer)
	}
}
