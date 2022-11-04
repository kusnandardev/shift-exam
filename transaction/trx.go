package transaction

import (
	"exam/products"
	"exam/utils"
	"fmt"
	"strings"
	"time"

	"github.com/rodaine/table"
)

type Receipt struct {
	ID       int
	DateTime time.Time
	Total    int
	Items    []Items
}

func NewTransaction(stock *[]products.Product, id int) Receipt {
	var trx Receipt
	var items []Items

trxItem:
	for {
		utils.Cls()
		switch transactionMenu(items) {
		case 1:
			items = append(items, addItem(stock))
		case 2:
			utils.Cls()
			products.CheckStock(stock)
			utils.Wait()
		case 3:
			payment(items)
			break trxItem
		default:
			fmt.Println("command not found of invalid input")
			utils.Wait()
		}
	}
	trx.ID = id + 1
	trx.Total = getGrandTotal(items)
	trx.Items = items
	trx.DateTime = time.Now()
	return trx
}

func payment(items []Items) {
	paid := 0
	total := getGrandTotal(items)

	utils.Cls()
	fmt.Println("Bill\t\t: ", total)
	fmt.Print("Paid\t\t: ")
	fmt.Scanf("%d\n", &paid)
	fmt.Println("Change\t\t: ", (paid - total))
	utils.Wait()
}

func addItem(stock *[]products.Product) Items {
	var item Items
	var id, qty int
	fmt.Print("input product id: ")
	fmt.Scanf("%d\n", &id)
	for i, v := range *stock {
		if v.Id == id {
			item.ItemName = v.Name
			fmt.Print("input quantity: ")
			fmt.Scanf("%d\n", &qty)
			item.Qty = qty
			item.SubTotal = v.Price * qty

			(*stock)[i].Substract(qty)
		}
	}

	return item
}

func transactionMenu(items []Items) int {
	selectedMenu := 0

	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("Transaction")
	fmt.Println(strings.Repeat("=", 60))
	if len(items) > 0 {
		printItems(items)
	}
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("1. add more item | 2. see product list | 3. payment")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Print("Pilih Menu: ")
	_, err := fmt.Scanf("%d\n", &selectedMenu)
	if err != nil {
		selectedMenu = 0
	}
	return selectedMenu
}

func printItems(items []Items) {
	tbl := table.New("No.", "Name", "Qty", "Sub Total")
	for i, item := range items {
		tbl.AddRow(i+1, item.ItemName, item.Qty, item.SubTotal)
	}
	tbl.Print()
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("Grand Total: ", getGrandTotal(items))
}

func getGrandTotal(items []Items) int {
	grandTotal := 0
	for _, item := range items {
		grandTotal += item.SubTotal
	}
	return grandTotal
}
