package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func readCSV(f io.Reader, separator rune) [][]string {
	r := csv.NewReader(f)

	r.Comma = separator

	l := [][]string{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		l = append(l, record)
	}
	return l
}

func main() {
	f, err := os.Open("country.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Println(readCSV(f, ','))

	f, err = os.Open("country.tsv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Println(readCSV(f, '\t'))
}
