package crudrun

import (
	"fmt"
	"model/postgres"
	"model/strac"
)

func CRUDProductRun(p *postgres.ProductRepo) {
	var choice int
	for {
		fmt.Println("0.Exit\t1.Create Product\t2.Update Product\t3.Delete Product\t4.Read Product\t5.Issue by id")
		fmt.Scan(&choice)
		switch choice {
		case 0:
			return
		case 1:
			var product strac.Product
			fmt.Printf("Isim kiriting: ")
			fmt.Scanln(&product.Name)
			fmt.Printf("Tavsifni kiriting: ")
			fmt.Scanln(&product.Description)
			fmt.Printf("Narxini kiriting: ")
			fmt.Scan(&product.Price)
			fmt.Printf("Stok miqdorini kiriting: ")
			fmt.Scan(&product.StockQuantity)

			err := p.CreateProduct(product)
			if err != nil {
				fmt.Println("Error creating product:", err)
				continue
			}
		case 2:
			var product strac.Product
			fmt.Printf("Qaysi id ni o`zgartirish kerak: ")
			fmt.Scan(&product.Id)

			fmt.Printf("Mahsulot nomini kiriting: ")
			fmt.Scanln(&product.Name)

			fmt.Printf("Tavsifni o`zgartiring: ")
			fmt.Scanln(&product.Description)

			fmt.Printf("Narxini o`zgartiring: ")
			fmt.Scan(&product.Price)

			err := p.UpdatePrpduct(product)
			if err != nil {
				fmt.Println("Error updating product:", err)
				continue 
			}
		case 3:
			var id int
			fmt.Printf("Qaysi id dagi malumotni o`chirmoqchisiz?: ")
			fmt.Scan(&id)
			err := p.DeleteProduct(id)
			if err != nil {
				fmt.Println("Error deleting product:", err)
				continue
			}
		case 4:
			products, err := p.GetAllProducts()
			if err != nil {
				fmt.Println("Error getting products:", err)
				continue
			}

			for _, k := range products {
				fmt.Println(k.Name, k.Price, k.Description, k.StockQuantity)
			}
		case 5:
			var id int
			fmt.Printf("Id kiriting: ")
			fmt.Scan(&id)
			product, err := p.GetProductById(id)
			if err != nil {
				fmt.Println("Error getting product:", err)
				continue
			}
			fmt.Println(product.Name, product.Description, product.Price, product.StockQuantity)
		default:
			fmt.Println("Noto'g'ri buyruq, qaytadan urinib ko'ring.")
		}
	}
}
