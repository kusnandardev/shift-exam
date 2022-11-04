package main

import (
	"exam/products"
	"exam/transaction"
	"exam/utils"
	"fmt"
	"strings"
)

func main() {

	prod := products.Product{}
	stock := prod.LoadProduct()
	trx := []transaction.Receipt{}

infinite:
	for {
		utils.Cls()
		switch menu() {
		case 1:
			trx = append(trx, transaction.NewTransaction(&stock, len(trx)))
		case 2:
			if len(trx) > 0 {
				transaction.GetTransactionReport(trx)
			} else {
				fmt.Println("No transaction was created")
				utils.Wait()
			}
		case 3:
			utils.Cls()
			products.CheckStock(&stock)
			utils.Wait()
		case 0:
			break infinite
		default:
			fmt.Println("command not found of invalid input")
			utils.Wait()
		}

	}
}

func menu() int {
	selectedMenu := 0
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("1. New Transaction")
	fmt.Println("2. Transaction Report")
	fmt.Println("3. List Product")
	fmt.Println("0. Exit")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Print("choose menu: ")
	_, err := fmt.Scanf("%d\n", &selectedMenu)
	if err != nil {
		selectedMenu = 99999999
	}
	return selectedMenu
}
