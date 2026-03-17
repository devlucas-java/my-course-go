package tax

import (
	"fmt"
	"testing"

	aluguel "github.com/devlucas/curso/test"
)

func TestAluguel(t *testing.T) {
	d := 1
	p := 200

	r := aluguel.Tiket{
		Day:   d,
		Price: p,
	}

	fmt.Println(r.Sum(r))

}
