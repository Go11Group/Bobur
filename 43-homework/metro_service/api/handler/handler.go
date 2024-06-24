package handler

import (
	"database/sql"
	"project/postgres"
)

type Handler struct {
	Card        *postgres.CardRepo
	Transaction *postgres.TransactionRepo
	Station     *postgres.StationRepo
	Terminal    *postgres.TerminalRepo
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		Card:        postgres.NewCardRepo(db),
		Transaction: postgres.NewTransactionRepo(db),
		Station:     postgres.NewStationRepo(db),
		Terminal:    postgres.NewTerminalRepo(db),
	}
}
