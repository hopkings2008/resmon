package parser

type Elem struct {
	Suffix string
	Index  int
}

type Config struct {
	// regex string to get elemets.
	Regstr string `json:regex`
}
