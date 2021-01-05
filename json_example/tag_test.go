package json_example_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type user struct {
	objectID  primitive.ObjectID `bson:"_id"`
	Name      *string            `json:"name" bson:"name"`
	CreatedAt *time.Time         `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt *time.Time         `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

func TestTag(t *testing.T) {
	input := `[
		{
			"name": "John",
			"created_at": "2020-12-28T12:20:15+09:00",
			"updated_at": "2020-12-28T15:20:15+09:00"
		},
		{
			"name": "Jane",
			"created_at": "2020-12-27T12:20:15+09:00"
		},
		{
			"name": "Jane",
			"created_at": "2020-12-27T12:20:15+09:00"
		},
		{
			"name": null,
			"created_at": null
		}
	]`

	var users []user
	if err := json.Unmarshal([]byte(input), &users); err != nil {
		t.Fatal("This shouldn't return error but got", err)
	}
	for _, u := range users {
		fmt.Printf("%v\n", u)
	}
}
