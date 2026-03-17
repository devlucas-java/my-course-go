package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	http.HandleFunc("/", findCepHandler)
	http.ListenAndServe(":3000", nil)
}

func findCepHandler(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte(`{"error": "not found"}`))
		res.WriteHeader(http.StatusNotFound)
	}

	cep := req.URL.Query().Get("cep")
	if cep == "" {
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte(`{"error": "cep is required"}`))
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	vc, err := findCep(cep)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(vc)
	c, err := json.Marshal(vc)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	if c != nil {
		res.Write(c)
	}
}

const firstUrl = "https://viacep.com.br/ws/"
const lastUrl = "/json/"

func findCep(cep string) (*ViaCep, error) {
	res, err := http.Get(firstUrl + cep + lastUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))

	var data ViaCep

	json.Unmarshal(body, &data)
	fmt.Println(data)
	return &data, nil

}
