package main

import "fmt"

func main() {
	defer fmt.Println("1 Primeira linha")
	fmt.Println("2 Segunda linha")
	defer fmt.Println("3 Terceira linha")
	fmt.Println("4 Quarta linha")
}
