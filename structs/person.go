package structs

type Person struct {
	Name     string
	Bio      string
	KnownFor []KnownFor
}

type KnownFor struct {
	Name string
	Link string
}
