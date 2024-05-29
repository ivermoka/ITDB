package lib


type OMDbResponse struct {
	Search       []Movie `json:"Search"`
	TotalResults string  `json:"totalResults"`
	Response     string  `json:"Response"`
}

type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}
