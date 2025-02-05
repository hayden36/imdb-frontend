package structs

type SearchResponse struct {
	Data  []SearchData `json:"d"`
	Query string       `json:"q"`
}
type SearchData struct {
	Image    ImageObject `json:"i"`
	ID       string      `json:"id"`
	Label    string      `json:"l"`
	Subtitle string      `json:"s"`
	Q        string      `json:"q"`
	Year     int         `json:"y"`
}
type ImageObject struct {
	Height   int    `json:"height"`
	ImageURL string `json:"imageUrl"`
	Width    int    `json:"width"`
}
