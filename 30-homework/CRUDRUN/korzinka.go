package crudrun

import (
	"fmt"
	"model/postgres"
	"model/strac"
)

func CRUDKorzinka1(k *postgres.KorzinkaRepo) {
	var korzinka1 strac.Korzinka
	var i int 
	for {
		fmt.Println("0.Exit\t1.Create Korzinka\t2.Update Korzinka\t3.Delete Korzinka\t4.Get all korzinka\t5.Get id korzinka")
		fmt.Scan(&i)
		switch i {
		case 0:
			return
		case 1:	
			fmt.Printf("User id kiriting: ")
			fmt.Scan(&korzinka1.User_id)
			fmt.Printf("Product id kiriting: ")
			fmt.Scan(&korzinka1.Product_id)

			err := k.CreateKorzinka(&korzinka1)
			if err != nil {
				panic(err)
			}
		case 2:
			var k1 strac.Korzinka
			fmt.Printf("Korzinka idsini kiriting: ")
			fmt.Scan(&k1.Id)
			fmt.Printf("Yangilanishi kerak bo`lgam user id: ")
			fmt.Scan(&k1.User_id)
			fmt.Printf("Yangilanishi kerak bo`lgan product id: ")
			fmt.Scan(&k1.Product_id)

			err := k.UpdateKorzinka(&k1)
			if err != nil {
				panic(err)
			}

		case 3:
			var id int 
			fmt.Printf("O`chirmoqchi bo`lgan id ni kiriting: ")
			fmt.Scan(&id)

			err := k.DeleteKorzinka(id)
			if err != nil {
				panic(err)
			}
		case 4:
			var korzinkas []strac.Korzinka

			korzinkas, err := k.GetAllKorzinka() 
			if err != nil {
				panic(err)
			}

			for _, v := range korzinkas {
				fmt.Println("User id: ", v.User_id, "Product id: ", v.Product_id)
			}

		case 5:
			var korzinka strac.Korzinka
			var id int
			fmt.Printf("Korzinka idsini kiriting: ")
			fmt.Scan(&id)

			korzinka, err := k.GetProductById(id)
			if err != nil {
				panic(err)
			}
			fmt.Println(korzinka)
		}
	}
}	