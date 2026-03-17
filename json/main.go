package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type Conta struct {
		Saldo int    `json:"saldo"`
		Name  string `json:"name"`
		Typee string `json:"type"`
	}

	account := Conta{Saldo: 3520, Name: "lucas", Typee: "corrente"}
	accountJson, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	fmt.Println(account)
	fmt.Println(string(accountJson))

	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		panic(err)
	}

	var accountX = Conta{}
	jsonX := []byte(`{"Saldo": 5782, "Name": "lucas2", "Typee": "account chained"}`)

	err = json.Unmarshal(jsonX, &accountX)

	if err != nil {
		panic(err)
	}

	fmt.Println(accountX)

	err = json.NewEncoder(os.Stderr).Encode(accountX)
	if err != nil {
		panic(err)
	}

}
