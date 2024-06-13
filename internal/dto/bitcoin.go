package dto

// название файла не оч понятное - почему биткоин
type TickerResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
