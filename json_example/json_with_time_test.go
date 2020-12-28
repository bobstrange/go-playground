package json_example_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/bobstrange/go-playground/json_example"
)

func TestMarshalJSONWithTime(t *testing.T) {
	u := json_example.User{
		ID:        "user01",
		Name:      "John Doe",
		Email:     "john.dow@email.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	b, err := json.Marshal(u)
	if err != nil {
		t.Fatal("It shouldn't get error but got ", err)
	}
	fmt.Printf("%s", b)
	// {"id":"user01","name":"John Doe","email":"john.dow@email.com","created_at":"2020-12-28T17:21:53.385044247+09:00","updated_at":"2020-12-28T17:21:53.385044327+09:00"}
}

func TestUnmarshalJSONWithTime(t *testing.T) {
	i := `
	{
		"id": "user01",
		"name": "John Doe",
		"email": "john.doe@email.com",
		"created_at": "2020-12-20T17:21:53+09:00",
		"updated_at": "2020-12-21T17:22:53.123+09:00"
	}
	`
	var u json_example.User
	err := json.Unmarshal([]byte(i), &u)
	if err != nil {
		t.Fatal("It shouldn't get error but got ", err)
	}
	fmt.Printf("%v", u)
}
