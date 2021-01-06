package zero_value_and_null_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

type User struct {
	Name      string     `json:"name"`
	Rings     *int       `json:"rings"`
	Nickname  *string    `json:"nickname,omitempty"`
	CreatedAt *time.Time `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type testCase struct {
	input string
	want  []*User
}

func String(v string) *string {
	return &v
}

func getTestCase() *testCase {
	tick, _ := time.Parse(time.RFC3339, "2020-01-05T01:23:45+09:00")

	return &testCase{
		input: `
		[
			{
				"name": "Micheal Jordan",
				"rings": 6,
				"created_at": "2020-01-05T01:23:45+09:00",
				"deleted_at": "2020-01-05T01:23:45+09:00"
			},
			{
				"name": "Lebron James",
				"nickname": "",
				"rings": 4,
				"created_at": "2020-01-05T01:23:45+09:00"
			},
			{
				"name": "Kyrie Irving",
				"nickname": null,
				"rings": 0,
				"created_at": "2020-01-05T01:23:45+09:00",
				"deleted_at": null
			},
			{
				"name": "Dummy",
				"created_at": "2020-01-01T00:00:00Z",
				"dummy_field": "test"
			}
		]
		`,
		want: []*User{
			{
				Name:      "Michel Jordan",
				Nickname:  String("MJ"),
				CreatedAt: &tick,
				DeletedAt: &tick,
			},
			{
				Name:      "Lebron James",
				Nickname:  String(""),
				CreatedAt: &tick,
				DeletedAt: nil,
			},
			{
				Name:      "Kyrie Irving",
				Nickname:  nil,
				CreatedAt: &tick,
				DeletedAt: nil,
			},
		},
	}
}

func parse(s string) ([]*User, error) {
	var users []*User
	err := json.Unmarshal([]byte(s), &users)
	return users, err
}

func TestMarshal(t *testing.T) {
	ex := getTestCase()
	users, err := parse(ex.input)
	if err != nil {
		t.Fatal("This shouldn't return any errors but got ", err)
	}
	for i, d := range users {
		log.Printf("i: %d name: %v\n", i, d.Name)
		log.Println("i: ", i, " nickname:", d.Nickname)
		log.Println("i: ", i, " deletedAt:", d.DeletedAt)
	}
}

func TestMarshalAndUnmarshal(t *testing.T) {
	ex := getTestCase()
	users, err := parse(ex.input)
	data, err := json.Marshal(users)
	if err != nil {
		log.Fatalln("Fail to marshal with error:", err)
	}
	log.Println(string(data))
}
