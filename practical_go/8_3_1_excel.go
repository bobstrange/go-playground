package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func writeToExcel() {
	out := excelize.NewFile()
	out.SetCellValue("Sheet1", "A1", "Hi there")
	if err := out.SaveAs("test.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func readFromExcel() {
	in, err := excelize.OpenFile("test.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	cell, err := in.GetCellValue("Sheet1", "A1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cell)
}

func main() {
	writeToExcel()
	readFromExcel()
}
