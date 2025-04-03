package types

type Stock struct {
	StockId int16  `json:"stockid"`
	Name    string `json:"name"`
	Price   int16  `json:"price"`
	Company string `json:"company"`
}

type Response struct {
	ID      int16  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}
