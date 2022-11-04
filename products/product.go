package products

type Product struct {
	Id    int
	Name  string
	Price int
	Stock int
}

func (p *Product) Substract(n int) {
	p.Stock -= n
}

func (p Product) LoadProduct() []Product {
	return []Product{
		{
			Id:    1,
			Name:  "Buku",
			Price: 10000,
			Stock: 30,
		},
		{
			Id:    2,
			Name:  "Pulpen",
			Price: 5000,
			Stock: 50,
		},
		{
			Id:    3,
			Name:  "Pensil",
			Price: 4000,
			Stock: 50,
		},
		{
			Id:    4,
			Name:  "Penghapus",
			Price: 2000,
			Stock: 40,
		},
		{
			Id:    5,
			Name:  "Penggaris",
			Price: 8000,
			Stock: 15,
		},
	}
}
