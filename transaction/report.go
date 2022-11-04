package transaction

import (
	"exam/utils"
	"fmt"
	"strings"
	"time"

	"github.com/rodaine/table"
)

func GetTransactionReport(list []Receipt) {
	utils.Cls()
	var selectedId int
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("Report Transaction")
	fmt.Println(strings.Repeat("=", 60))
	tbl := table.New("ID", "Date", "Total")
	for _, item := range list {
		tbl.AddRow(item.ID, item.DateTime.Format(time.RFC822), item.Total)
	}
	tbl.Print()
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("Input transaction ID to get Detail | press enter to go back")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Print("Detail Transaction ID:")
	fmt.Scanf("%d\n", &selectedId)
	for _, v := range list {
		if v.ID == selectedId {
			fmt.Println(strings.Repeat("=", 60))
			printItems(v.Items)
			fmt.Println(strings.Repeat("=", 60))
			utils.Wait()
			GetTransactionReport(list)
		}
	}
}
