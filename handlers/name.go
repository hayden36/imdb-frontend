package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/hayden36/imdb-frontend/pages"
	"github.com/hayden36/imdb-frontend/structs"
	"github.com/hayden36/imdb-frontend/utils"
)

func HandleNameRoute() func(c *gin.Context) {
	return func(c *gin.Context) {
		nameParameter := c.Param("name")
		req, _ := http.NewRequest("GET", "https://www.imdb.com/name/"+nameParameter, nil)
		req.Header.Set("User-Agent", utils.USER_AGENT)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			pages.Error(500, fmt.Sprintf("Error sending request: %s", err)).Render(c, c.Writer)
			log.Printf("Error sending request to IMDB: %s", err)
			return
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			if res.StatusCode == 404 {
				log.Printf("404")
				pages.Error(404, "Page not found.").Render(c, c.Writer)
				return
			}
			pages.Error(500, fmt.Sprintf("Status code was not 200.\n Got a %s.", res.Status)).Render(c, c.Writer)
			log.Printf("status code error: %d %s", res.StatusCode, res.Status)
			return
		}
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		nameSelection := doc.Find(".hero__primary-text").Text()
		bioSelection := doc.Find(".ipc-html-content-inner-div").First()
		var knownForSlice []structs.Link
		bioHtml, _ := bioSelection.Html()

		personImage := doc.Find(".ipc-poster").Find("img.ipc-image").AttrOr("src", "")
		doc.Find("div.ipc-metadata-list-summary-item__tc").Each(func(i int, s *goquery.Selection) {
			title := s.Find("a.ipc-metadata-list-summary-item__t").Text()
			link := s.Find("a").AttrOr("href", "none")
			tempKnownFor := structs.Link{Name: title, Link: link}
			knownForSlice = append(knownForSlice, tempKnownFor)
		})
		// Personal Details
		var personalDetails structs.PersonalDetails
		personalDetailsSection := doc.Find("[data-testid=PersonalDetails]").Find("[data-testid=nm_pd_hd]")

		// Oficial Sites
		personalDetailsSection.Find("[data-testid=details-officialsites]").First().Find(".ipc-inline-list__item").Each(func(i int, s *goquery.Selection) {
			title := s.Find("a").Text()
			link := s.Find("a").AttrOr("href", "none")
			personalDetails.OfficialSites = append(personalDetails.OfficialSites, structs.Link{Name: title, Link: link})
		})
		person := structs.Person{Name: nameSelection, Bio: bioHtml, KnownFor: knownForSlice, PersonalDetails: personalDetails, Image: personImage}

		pages.NameView(person).Render(c, c.Writer)
	}
}
