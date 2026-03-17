package main

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         int
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

type Category struct {
	ID   int
	Name string
}

func main() {

	dsn := "root:root@tcp(localhost:3306)/banco?parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	products1 := []Product{
		{Name: "teste", Price: 999.0, CategoryID: 2}}
	// 	{Name: "carro", Price: 3423.0, CategoryID: 1},
	// 	{Name: "moto2", Price: 3423.0, CategoryID: 1},
	// 	{Name: "carro2", Price: 3423.0, CategoryID: 1},
	// }

	db.Create(products1)
	var product Product
	db.Delete(&product, 9)
	// db.Create(products)

	products := []Product{}
	db.Preload("Category").Find(&products)

	for _, p := range products {
		fmt.Fprintln(os.Stderr, p)
	}
	fmt.Println("-------------------------------")
	var product2 Product
	db.Preload("Category").First(&product2, 8)
	fmt.Println(product2)

}
