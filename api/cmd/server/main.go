package main

import (
	"fmt"
	"net/http"

	"github.com/devlucas-java/curso/api/config"
	"github.com/devlucas-java/curso/api/internal/entity"
	"github.com/devlucas-java/curso/api/internal/infra/database"
	"github.com/devlucas-java/curso/api/internal/webserver/handlers"
	"github.com/go-chi/chi"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println(err)
	}
	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	c := chi.NewRouter()

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)
	c.Post("/product", productHandler.CreateProduct)
	c.Get("/product/{id}", productHandler.FindById)
	c.Delete("/product/{id}", productHandler.DeleteProduct)
	c.Get("/product/find", productHandler.FindAllProducts)
	c.Patch("/product", productHandler.UpdateProduct)

	http.ListenAndServe(":8080", c)
}
