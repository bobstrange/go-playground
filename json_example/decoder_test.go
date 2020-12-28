package json_example_test

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestDecoder(t *testing.T) {
	const jsonStream = `
	{"Name": "Ed", "Text": "Knock knock."}
	{"Name": "Sam", "Text": "Who's there."}
	{"Name": "Ed", "Text": "Go fmt."}
	{"Name": "Sam", "Text": "Go fmt who ?"}
	{"Name": "Ed", "Text": "Go fmt yourself!"}
`
	type Message struct {
		Name string
		Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			t.Fatal("This shouldn't return error but got", err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}
