package postgres

import (
	"database/sql"
	"project/model"

	"github.com/google/uuid"
)

type CardRepo struct {
	Db *sql.DB
}

func NewCardRepo(db *sql.DB) *CardRepo {
	return &CardRepo{db}
}

// Card create qilish
func (c *CardRepo) CreateCard(card *model.Card) error {
	_, err := c.Db.Exec("insert into card(id, number, user_id) values($1, $2, $3)", uuid.NewString(), card.Number, card.UserID)
	if err != nil {
		return err
	}

	return nil
}

// Card ni idisini delete qilish
func (c *CardRepo) DeleteCard(id string) error {
	_, err := c.Db.Exec("delete from card where id=$1", id)
	if err != nil {
		return err
	}

	return nil 
}


// Card id si orqali Update qilish
func (c *CardRepo) UpdateCard(id string, card *model.Card) (*model.Card, error) {
	_, err := c.Db.Exec("update card set number=$1, user_id=$2 where id=$3", card.Number, card.UserID, id)
	if err != nil {
		return nil, err
	}
	res, err := c.GetCardId(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Card ni id si orqali chiqarish
func (c *CardRepo) GetCardId(id string) (*model.Card, error) {
	card := model.Card{}
	row := c.Db.QueryRow("select id, number, user_id from card")
	err := row.Scan(&card.Id, &card.Number, &card.UserID)
	if err != nil {
		return nil, err
	}

	return &card, nil
}

// Card ni hamma ma'lumotini chiqarish
func (c *CardRepo) GetCardAll() ([]model.Card, error) {

	rows, err := c.Db.Query("select * from card")
	if err != nil {
		return nil, err
	}
	cards := []model.Card{}

	for rows.Next() {
		card := model.Card{}
		err = rows.Scan(&card.Id, &card.Number, &card.UserID)
		if err != nil {
			return nil, err
		}

		cards = append(cards, card)

	}

	return cards, nil
}
