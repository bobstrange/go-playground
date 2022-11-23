package main

import (
	"log"

	"github.com/xuri/excelize/v2"
)

type Country struct {
	Name       string
	ISOCode    string
	Population int
}

func (c Country) HeaderColumns() []interface{} {
	return []interface{}{"国名", "ISOコード", "人口"}
}

func (c Country) Columns() []interface{} {
	return []interface{}{c.Name, c.ISOCode, c.Population}
}

func main() {
	lines := []Country{
		{"アメリカ合衆国", "US", 328_239_523},
		{"日本", "JP", 126_000_000},
		{"インド", "IN", 1_380_004_385},
	}
	f := excelize.NewFile()
	sw, err := f.NewStreamWriter("Sheet1")
	if err != nil {
		log.Fatal(err)
	}

	for i, line := range lines {
		if i == 0 {
			cell, _ := excelize.CoordinatesToCellName(1, i+1)
			sw.SetRow(cell, line.HeaderColumns())
		}
		cell, _ := excelize.CoordinatesToCellName(1, i+2)
		sw.SetRow(cell, line.Columns())
	}

	if err := sw.Flush(); err != nil {
		log.Fatal(err)
	}

	if err := f.SaveAs("countries.xlsx"); err != nil {
		log.Fatal(err)
	}
}
