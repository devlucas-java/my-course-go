package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

// huffman, aritmetica e algoritimos lepel, ziv

type Product struct {
	ID        string
	Name      string
	Price     float64
	CreatedAt time.Time
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:        uuid.New().String(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
}

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/banco?parseTime=true")

	if err != nil {
		fmt.Fprint(os.Stderr, err)
		panic(err)
	}
	defer db.Close()

	// insert
	product := NewProduct("moto", 5826)
	err = insertProduct(db, *product)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		panic(err)
	}
	product2 := NewProduct("rifle", 3824)
	err = insertProduct(db, *product2)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		panic(err)
	}

	products, err := selectAllProducts(db)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		panic(err)
	}

	for _, p := range products {
		fmt.Fprintf(os.Stderr, "%+v\n", p)
	}

	product.Price = 1234
	product.ID = "d0d70899-1980-ale2-efc1532b8ela"
	product.Name = "moto 2"

	// update
	err = updateProducts(db, *product)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		panic(err)
	}
}

func insertProduct(db *sql.DB, product Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price, created_at) values(?, ?, ?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price, product.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func updateProducts(db *sql.DB, product Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)

	if err != nil {
		return err
	}

	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select * from products where id = ?")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	var product Product
	err = stmt.QueryRow(id).Scan(&product.Name, &product.Price, &product.CreatedAt)
	// err = stmt.QueryRowContext(context.Background(), id).Scan(&product.Name, &product.Price, &product.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select * from products")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	for _, p := range products {
		fmt.Fprintf(os.Stderr, "%+v\n", p)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}
