package postgres

import (
	"database/sql"
	"project/model"

	"github.com/google/uuid"
)

type TerminalRepo struct {
	Db *sql.DB
}

func NewTerminalRepo(db *sql.DB) *TerminalRepo {
	return &TerminalRepo{db}
}

//  Terminal ni create qilish 
func (t *TerminalRepo) CreateT(tr *model.Terminal) error {
	_, err := t.Db.Exec("insert into terminal(id, station_id) values($1, $2,)", uuid.NewString(), tr.Id, tr.StationID)
	if err != nil {
		return err
	}

	return nil
}

// Terminal ni idisini delete qilish
func (t *TerminalRepo) DeleteT(id string) error {
	_, err := t.Db.Exec("delete from terminal where id=$1", id)
	if err != nil {
		return err
	}

	return nil
}


// Terminal id si orqali Update qilish
func (t *TerminalRepo) UpdateT(id string, tr *model.Terminal) (*model.Terminal, error) {
	_, err := t.Db.Exec("update station station=$1 where id=$2", tr.StationID, id)
	if err != nil {
		return nil, err
	}
	res, err := t.GetTId(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Terminal ni id si orqali chiqarish
func (t *TerminalRepo) GetTId(id string) (*model.Terminal, error) {
	st := model.Terminal{}
	row := t.Db.QueryRow("select id, station_id from terminal")
	err := row.Scan(&st.Id, &st.StationID)
	if err != nil {
		return nil, err
	}

	return &st, nil
}

// Terminal ni hamma ma'lumotini chiqarish
func (t *TerminalRepo) GetTAll() ([]model.Terminal, error) {

	rows, err := t.Db.Query("select * from terminal")
	if err != nil {
		return nil, err
	}
	trs := []model.Terminal{}

	for rows.Next() {
		tr := model.Terminal{}
		err = rows.Scan(&tr.Id, &tr.StationID)
		if err != nil {
			return nil, err
		}

		trs = append(trs, tr)

	}

	return trs, nil
}
