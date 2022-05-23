package model

//TuitRequest
type TuitRequest struct {
	UserId  string `json:"userId"`
	Content string `json:"content"`
}

type Platuit struct {
	ID      string `json:"id"`
	UserId  string `json:"userId"`
	Content string `json:"content"`
}

type GetPlatuitsResponse []Platuit
