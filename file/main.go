package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	oss, err := os.Create("teste.txt")

	if err != nil {
		panic(err)
	}

	size, err := oss.Write([]byte("teste 123"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("seu tamanho e \b", size)
	oss.Close()

	file, err := os.Open("teste.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1)

	for {
		f, err := reader.Read(buffer)
		if err != nil {
			fmt.Println("erro in loop")
			break
		}
		fmt.Println(string(buffer[:f]))
	}

	file.Close()

}
