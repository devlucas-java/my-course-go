package aluguel

type Tiket struct {
	Day   int
	Price int
}

func (t Tiket) Sum(p Tiket) int {
	retult := p.Price * p.Day
	if retult > 300 {
		return retult - 50
	}
	return retult
}
