package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const firstUrl = "https://viacep.com.br/ws/"
const lastUrl = "/json/"

type cep struct {
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
	for _, arg := range os.Args[1:] {
		res, err := http.Get(firstUrl + arg + lastUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v", err)
			continue
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v", err)
			continue
		}

		var viaCep cep
		err = json.Unmarshal(body, &viaCep)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v", err)
			continue
		}

		file, err := os.Create("city.json")
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer file.Close()

		err = json.NewEncoder(file).Encode(viaCep)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Print(arg)
		defer res.Body.Close()
	}
}
