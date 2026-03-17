package main

import (
	"fmt"
	"os"
)

func main() {

	for _, ag := range os.Args[1:] {
		fmt.Println(ag)
	}

}
