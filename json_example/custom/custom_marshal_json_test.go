package custom_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

type animal int

const (
	Unknown animal = iota
	Gopher
	Zebra
)

// UnmarshalJSON で []byte から求めている型に変換する処理を定義できる
func (a *animal) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	default:
		*a = Unknown
	case "gopher":
		*a = Gopher
	case "zebra":
		*a = Zebra
	}
	return nil
}

// MarshalJSON で、型から []byte への変換する処理を定義できる
func (a animal) MarshalJSON() ([]byte, error) {
	var s string
	switch a {
	default:
		s = "unknown"
	case Gopher:
		s = "gopher"
	case Zebra:
		s = "zebra"
	}
	return json.Marshal(s)
}

func TestUnmarshalAnimal(t *testing.T) {
	blob := `["gopher","armadillo","zebra","unknown","gopher","bee","gopher","zebra"]`
	var zoo []animal
	if err := json.Unmarshal([]byte(blob), &zoo); err != nil {
		t.Fatal("This shouldn't return error but got", err)
	}
	mapping := make(map[animal]int)

	for _, animal := range zoo {
		mapping[animal] += 1
	}
	fmt.Printf("Gophers: %d\nZebras: %d\nUnknown: %d\n", mapping[Gopher], mapping[Zebra], mapping[Unknown])
}
