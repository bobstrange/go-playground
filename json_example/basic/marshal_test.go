package basic_test

import (
	"encoding/json"
	"os"
	"testing"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func TestMarshal(t *testing.T) {
	group := ColorGroup{
		ID:   1,
		Name: "Reds",
		Colors: []string{
			"Crimson",
			"Red",
			"Ruby",
			"Maroon",
		},
	}
	b, err := json.Marshal(group)
	if err != nil {
		t.Fatal("This shouldn't return error but got", err)
	}
	os.Stdout.Write(b)
}
