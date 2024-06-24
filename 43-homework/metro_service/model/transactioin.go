package model

type Transaction struct {
	Id              string `json:"id"`
	CardID          string `json:"card_id"`
	Amount          int    `json:"amount"`
	TerminalID     string `json:"terminal_id"`
	TransactionType string `json:"transaction_type"`
}
