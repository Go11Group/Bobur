package postgres

import "database/sql"


type TransportRepo struct {
	Db *sql.DB
}

func NewTransportRepo(db *sql.DB) *TransportRepo {
	return &TransportRepo{db}
}

