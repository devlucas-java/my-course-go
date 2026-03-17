package main

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

type Product struct {
	ID         int
	Name       string
	Price      float64
	CategoryID int
	Category   Category
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

	products := []Product{}

	fmt.Println("-------------------")
	db.Limit(2).Offset(0).Find(&products)
	for _, p := range products {
		fmt.Fprintln(os.Stderr, p)
	}

	fmt.Println("-------------------")
	products2 := []Product{}
	//db.Where("price > ?", 40).Find(&products2)
	db.Where("name LIKE ?", "%book%").Find(&products2)
	for _, p := range products2 {
		fmt.Fprintln(os.Stderr, p)
	}

	fmt.Println("-------------------")

	fmt.Println("print ok ")
}
