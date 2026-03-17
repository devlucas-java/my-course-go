package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/devlucas-java/curso/api/internal/dto"
	"github.com/devlucas-java/curso/api/internal/entity"
	"github.com/devlucas-java/curso/api/internal/infra/database"
	"github.com/go-chi/chi"
)

type ProductHandler struct {
	DB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{DB: db}
}

func (ph *ProductHandler) CreateProduct(res http.ResponseWriter, req *http.Request) {
	var product dto.ProductDTO
	err := json.NewDecoder(req.Body).Decode(&product)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}

	p := entity.NewProduct(product.Name, product.Price)
	err = ph.DB.CreateProduct(p)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(res).Encode(p)
	res.WriteHeader(http.StatusCreated)
}

func (ph *ProductHandler) UpdateProduct(res http.ResponseWriter, req *http.Request) {

	var product dto.ProductDTO
	err := json.NewDecoder(req.Body).Decode(&product)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Erro in decoder json"))
	}

	p := entity.NewProduct(product.Name, product.Price)

	err = ph.DB.UpdateProduct(p)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
	}

	res.Write([]byte("Product updated"))
	res.WriteHeader(http.StatusOK)
}

func (ph *ProductHandler) FindById(w http.ResponseWriter, r *http.Request) {
	// Pegar o id da rota
	id := chi.URLParam(r, "id")

	// Buscar o produto
	product, err := ph.DB.FindById(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Product not found"})
		return
	}

	// Resposta com sucesso
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		// Caso ocorra erro na codificação JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
}

func (ph *ProductHandler) DeleteProduct(res http.ResponseWriter, req *http.Request) {

	id := chi.URLParam(req, "id")
	if id == "" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Id is required"))
	}
	err := ph.DB.DeleteProduct(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Error in delete product"))
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Product deleted"))
}

func (ph *ProductHandler) FindAllProducts(res http.ResponseWriter, req *http.Request) {

	page := chi.URLParam(req, "page")
	limit := chi.URLParam(req, "limit")
	sort := chi.URLParam(req, "sort")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Error in convert page to int"))
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Error in convert limit to int"))
	}

	products, err := ph.DB.FindAllProducts(pageInt, limitInt, sort)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Error in find all products"))
	}
	json.NewEncoder(res).Encode(products)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

}
