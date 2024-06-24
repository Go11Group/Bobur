package postgres

import (
	"database/sql"
	"project/model"
)

type StationRepo struct {
	Db *sql.DB
}

func NewStationRepo(db *sql.DB) *StationRepo {
	return &StationRepo{db}
}

func (s *StationRepo) STCreate(station *model.Station) error {
	_, err := s.Db.Exec("insert into station(name)values ($1)", station.Name)
	return err
}

// Station ni idisini delete qilish
func (s *StationRepo) STDelete(id string) error {
	_, err := s.Db.Exec("delete from station where id=$1", id)
	if err != nil {
		return err
	}

	return nil
}

// Station id si orqali Update qilish
func (s *StationRepo) STUpdate(id string, st *model.Station) (*model.Station, error) {
	_, err := s.Db.Exec("update station name=$1 where id=$2", st.Name, id)
	if err != nil {
		return nil, err
	}
	res, err := s.STGetId(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Station ni id si orqali chiqarish
func (s *StationRepo) STGetId(id string) (*model.Station, error) {
	st := model.Station{}
	row := s.Db.QueryRow("select id, name from station")
	err := row.Scan(&st.Id, &st.Name)
	if err != nil {
		return nil, err
	}

	return &st, nil
}

// Station ni hamma ma'lumotini chiqarish
func (s *StationRepo) STGetAll() ([]model.Station, error) {

	rows, err := s.Db.Query("select * from card")
	if err != nil {
		return nil, err
	}
	sts := []model.Station{}

	for rows.Next() {
		st := model.Station{}
		err = rows.Scan(&st.Id, &st.Name)
		if err != nil {
			return nil, err
		}

		sts = append(sts, st)

	}

	return sts, nil
}
