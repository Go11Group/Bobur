package postgres

import (
	"database/sql"
	"project/model"

	"github.com/google/uuid"
)


type TransactionRepo struct {
	Db *sql.DB
}

func NewTransactionRepo(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{db}
}

func (t *TransactionRepo) CreateTR(tr *model.Transaction) error {
	_, err := t.Db.Exec("insert into transaction(id, card_id, amount, terminal_id, transaction_type) values($1, $2, $3, $4, $5)",
						uuid.NewString(), tr.CardID, tr.Amount, tr.TerminalID, tr.TransactionType)
	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionRepo) DeleteTR(id string) error {
	_, err := t.Db.Exec("delete from transaction where id=$1", id)
	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionRepo) UpdateTR(id string, tr *model.Transaction) (*model.Transaction, error) {
	_, err := t.Db.Exec("update transaction set card_id=$1, amount=$2, terminal_id=$3, transaction_type=$4 where id=$5", 
						tr.CardID, tr.Amount, tr.TerminalID, tr.TransactionType, id)
	
	if err != nil {
		return nil, err
	}
	res, err := t.GetTRByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil 
}

func (t *TransactionRepo) GetTRByID(id string) (*model.Transaction, error) {
	tr := model.Transaction{}

	row := t.Db.QueryRow("select id, card_id, amount, terminal_id, transaction_type from transaction")
	err := row.Scan(&tr.Id, &tr.CardID, &tr.Amount, &tr.TerminalID, &tr.TransactionType)
	
	if err != nil {
		return nil, err
	}

	return &tr, nil
}

func (t *TransactionRepo) GetTRAll() ([]model.Transaction, error) {
	rows, err := t.Db.Query("select id, card_id, amount, terminal_id, transaction_type from transaction")
	if err != nil {
		return nil,  err
	}

	trs := []model.Transaction{}

	for rows.Next() {
		tr := model.Transaction{}

		err = rows.Scan(&tr.Id, &tr.CardID, &tr.Amount, &tr.TerminalID, &tr.TransactionType)
		if err != nil {
			return nil, err
		}
		trs = append(trs, tr)
	}

	return trs, nil
}

