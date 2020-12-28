package json_example_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestValid(t *testing.T) {
	good := `{"example": 1}`
	bad := `{"example":2:]}}`

	fmt.Println(
		json.Valid([]byte(good)),
		json.Valid([]byte(bad)),
	)
}
