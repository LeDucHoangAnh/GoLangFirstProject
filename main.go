package main

import (
	"fmt"

	"go-module/c4f"

	"github.com/leekchan/accounting"
)

func main() {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	fmt.Println(ac.FormatMoney(123456789.213123))
	c4f.Println("this is red color")
}
