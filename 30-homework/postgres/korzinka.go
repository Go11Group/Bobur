package postgres

import (
	"database/sql"
	"fmt"
	"model/strac"
)

type KorzinkaRepo struct {
	DB *sql.DB
}

func NewKorzinkaRepo(DB *sql.DB) *KorzinkaRepo {
	return &KorzinkaRepo{DB}
}

func (k *KorzinkaRepo) CreateKorzinka(korzinka *strac.Korzinka) error {
	tr, err := k.DB.Begin()
	if err != nil {
		return fmt.Errorf("tranzaktsiyani boshlashda xato: %w", err)
	}

	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()

	_, err = tr.Exec("insert into korzinka(user_id, product_id) values($1, $2)",
		korzinka.User_id, korzinka.Product_id)
	if err != nil {
		panic(err)
	}

	return nil
}

func (k *KorzinkaRepo) UpdateKorzinka(korzinka *strac.Korzinka) error {
	tr, err := k.DB.Begin()	
	if err != nil {
		return fmt.Errorf("tranzaktsiyani boshlashda xato: %w", err)
	}

	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()

	_, err = tr.Exec("update korzinka set product_id=$1 where product_id=$2", korzinka.Id, korzinka.Product_id)

	return nil
}

func (k *KorzinkaRepo) DeleteKorzinka(id int) error {
	tr, err := k.DB.Begin()
	if err != nil {
		return fmt.Errorf("tranzaktsiyani boshlashda xato: %w", err)
	}

	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()

	_, err = tr.Exec("delete from korzinka where id=$1", id)

	return nil
}

func (k *KorzinkaRepo) GetAllKorzinka() ([]strac.Korzinka, error) {
	tr, err := k.DB.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tr.Query("select user_id, product_id from korzinka")
	if err != nil {
		tr.Rollback()
		return nil, err
	}
	defer rows.Close()

	var korzinkas []strac.Korzinka
	for rows.Next() {
		var korzinka strac.Korzinka
		if err := rows.Scan(&korzinka.User_id, &korzinka.Product_id); err != nil {
			tr.Rollback()
			return nil, err
		}
		korzinkas = append(korzinkas, korzinka)
	}

	if err := rows.Err(); err != nil {
		tr.Rollback()
		return nil, err
	}

	return korzinkas, tr.Commit()

}


func (k *KorzinkaRepo) GetProductById(id int) (strac.Korzinka, error) {
	var korzinka strac.Korzinka

	err := k.DB.QueryRow("SELECT user_id, product_id FROM  WHERE id = $1", id).Scan(&korzinka.User_id, &korzinka.Product_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return korzinka, fmt.Errorf("mahsulot topilmadi, id: %d", id)
		}
		return korzinka, fmt.Errorf("mahsulotni olishda xato: %w", err)
	}

	return korzinka, nil
}
