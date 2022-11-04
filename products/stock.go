package products

import (
	"fmt"
	"strings"

	"github.com/rodaine/table"
)

func CheckStock(stock *[]Product) {
	tbl := table.New("ID", "Name", "Price", "Stock")
	for _, item := range *stock {
		tbl.AddRow(item.Id, item.Name, item.Price, item.Stock)
	}
	fmt.Println(strings.Repeat("=", 60))
	tbl.Print()
	fmt.Println(strings.Repeat("=", 60))

}
