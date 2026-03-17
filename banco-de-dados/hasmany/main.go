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
	ID       int
	Name     string
	Products []Product
}

func main() {

	dsn := "root:root@tcp(localhost:3306)/banco?parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	var products []Product
	db.Find(&products)

	for _, p := range products {
		fmt.Fprintln(os.Stderr, "--: ", p)
	}

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error

	for _, c := range categories {
		fmt.Fprintf(os.Stderr, "=============%+v============= \n", c.Name)
		for _, p := range c.Products {
			fmt.Fprintln(os.Stderr, "---:", p)
		}
	}

}
