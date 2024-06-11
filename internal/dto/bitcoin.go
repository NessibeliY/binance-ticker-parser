package dto

type TickerResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}