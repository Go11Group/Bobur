package crudrun

import (
	"fmt"
	"model/postgres"
	"model/strac"
)

func CRUDUserRun(u *postgres.UserRepo) {
	var a int
	for {
		fmt.Println("0.Exit\t1.Create User\t2.Update User\t3.Delete User\t4.Read User\t5.Issue by id")
		fmt.Scan(&a)
		switch a {
		case 0:
			return
		case 1:
			var user strac.User
			fmt.Printf("Isim kiriting: ")
			fmt.Scan(&user.Name)
			fmt.Printf("Email kiriting: ")
			fmt.Scan(&user.Email)
			fmt.Printf("Passowrd kiriting: ")
			fmt.Scan(&user.Password)

			err := u.CreateUser(user)
			if err != nil {
				panic(err)
			}
		case 2:
			var user strac.User
			fmt.Printf("Qaysi id ni o`zgartir kerak: ")
			fmt.Scan(&user.Id)

			fmt.Printf("Isim kiriting: ")
			fmt.Scan(&user.Name)

			fmt.Printf("Email kiriting: ")
			fmt.Scan(&user.Email)

			fmt.Println("Password kiriting: ")
			fmt.Scan(&user.Password)

			err := u.UpdateUser(user)
			if err != nil {
				panic(err)
			}

		case 3:
			var id int
			fmt.Printf("Qaysi id dagi malumotni o`chirmoqchisiz?: ")
			fmt.Scan(&id)
			err := u.DeleteUser(id)
			if err != nil {
				panic(err)
			}
		case 4:
			var users []strac.User

			users, err := u.GetAllUser()
			if err != nil {
				panic(err)
			}

			for _, k := range users {
				fmt.Println(k.Name, k.Email, k.Password)
			}
		case 5:
			var user strac.User
			var id int
			fmt.Printf("Id kiriting: ")
			fmt.Scan(&id)
			user, err := u.GetId(id)
			if err != nil {
				panic(err)
			}
			fmt.Println(user.Name, user.Email, user.Password)
		}

	}
}
