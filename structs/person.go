package structs

type Person struct {
	Name            string
	Bio             string
	KnownFor        []Link
	PersonalDetails PersonalDetails
	Image           string
}

type Link struct {
	Name string
	Link string
}

type PersonalDetails struct {
	OfficialSites []Link
	AlternateName string
	Height        string
	Born          string
	Parents       []string
	OtherWorks    string
	Spouse        string
	Children      string
	Relatives     string
}
