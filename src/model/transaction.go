package model

type Transaction struct {
	TotalValue float64 `json:"totalValue"`
	PayerID    string  `json:"payerID"`
	PayeeID    string  `json:"payeeID"`
	Items      []Item  `json:"items"`
}
