package structs

type MovieShow struct {
	Name    string
	Plot    string
	TopCast []Cast
}

type Cast struct {
	Name string
	Link string
}
