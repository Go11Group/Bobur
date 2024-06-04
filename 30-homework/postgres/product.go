package postgres

import (
	"database/sql"
	"fmt"
	"model/strac"
)

type ProductRepo struct {
	DB *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{db}
}

func (u *ProductRepo) CreateProduct(product strac.Product) error {
	tr, err := u.DB.Begin()
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

	_, err = tr.Exec("insert into products(name, description, price, stock_quantity) values($1, $2, $3, $4)",
		product.Name, product.Description, product.Price, product.StockQuantity)
	if err != nil {
		return fmt.Errorf("mahsulotni kiritishda xato: %w", err)
	}

	return nil
}

func (p *ProductRepo) UpdatePrpduct(product strac.Product) error {
	tr, err := p.DB.Begin()
	if err != nil {
		panic(err)
	}

	_, err = tr.Exec("update products set name=$1, description=$2, price=$3 where id = $4", product.Name, product.Description, product.Price, product.Id)
	if err != nil {
		tr.Rollback()
		return err
	}
	return tr.Commit()
}

func (u *ProductRepo) DeleteProduct(id int) error {
	tr, err := u.DB.Begin()
	if err != nil {
		panic(err)
	}

	_, err = tr.Exec("delete from products where id=$1", id)
	if err != nil {
		tr.Rollback()
		return err
	}
	return tr.Commit()
}

func (u *ProductRepo) GetAllProducts() ([]strac.Product, error) {
	tr, err := u.DB.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tr.Query("select name, description , price, stock_quantity from products")
	if err != nil {
		tr.Rollback()
		return nil, err
	}
	defer rows.Close()

	var products []strac.Product
	for rows.Next() {
		var product strac.Product
		if err := rows.Scan(&product.Name, &product.Description, &product.Price, &product.StockQuantity); err != nil {
			tr.Rollback()
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		tr.Rollback()
		return nil, err
	}

	return products, tr.Commit()
}

func (u *ProductRepo) GetProductById(id int) (strac.Product, error) {
	var product strac.Product

	err := u.DB.QueryRow("SELECT id, name, description, price, stock_quantity FROM products WHERE id = $1", id).Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.StockQuantity)
	if err != nil {
		if err == sql.ErrNoRows {
			return product, fmt.Errorf("mahsulot topilmadi, id: %d", id)
		}
		return product, fmt.Errorf("mahsulotni olishda xato: %w", err)
	}

	return product, nil
}
