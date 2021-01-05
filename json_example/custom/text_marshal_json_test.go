package custom_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

type size int

const (
	unrecognized size = iota
	small
	large
)

func (s *size) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {
	default:
		*s = unrecognized
	case "small":
		*s = small
	case "large":
		*s = large
	}
	return nil
}

func (s size) MarshalText() ([]byte, error) {
	var name string
	switch s {
	default:
		name = "unrecognized"
	case small:
		name = "small"
	case large:
		name = "large"
	}
	return []byte(name), nil
}

func TestUnmarshalText(t *testing.T) {
	blob := `["small","regular","large","unrecognized","small","normal","small","large"]`
	var inventory []size
	if err := json.Unmarshal([]byte(blob), &inventory); err != nil {
		t.Fatal("This shouldn't return error but got", err)
	}

	counts := make(map[size]int)
	for _, size := range inventory {
		counts[size] += 1
	}
	fmt.Printf("Inventory counts:\n*small: %d\n*large: %d\n*unrecognized: %d\n", counts[small], counts[large], counts[unrecognized])

}
