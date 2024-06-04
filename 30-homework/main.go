package main

import (
	"fmt"
	crudrun "model/CRUDRUN"
	"model/postgres"
)

func main() {
	db, err := postgres.ConnDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	userdb := postgres.NewUserRepo(db)
	product := postgres.NewProductRepo(db)
	korzinka1 := postgres.NewKorzinkaRepo(db)

	var a int

	for {
		fmt.Println("0.Exit\t1.User\t2.Product\t3.Korzinka")
		fmt.Scan(&a)
		switch a {
		case 0:
			fmt.Println("Successfully connected!")
			return
		case 1:
			crudrun.CRUDUserRun(userdb)
		case 2:
			crudrun.CRUDProductRun(product)
		case 3:
			crudrun.CRUDKorzinka1(korzinka1)
		}
	}

}
