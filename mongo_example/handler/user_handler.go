package handler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/bobstrange/go-playground/mongo_example/entities"
	"github.com/bobstrange/go-playground/mongo_example/repositories"
	"go.mongodb.org/mongo-driver/bson"
)

func Handle(repo repositories.RepoIface) ([]byte, error) {
	input := `
	[
		{
			"id": "test_user_04",
			"first_name": "Lebron",
			"last_name": "James",
			"nickname": "",
			"age": 0,
			"created_at": "2021-01-05T12:34:50.12345+09:00"
		},
		{
			"id": "test_user_05",
			"first_name": "James",
			"last_name": "Harden",
			"nickname": "The Beard",
			"age": 31,
			"created_at": "2021-01-03T12:34:50+09:00",
			"updated_at": "2021-01-05T10:11:20Z"
		},
		{
			"id": "test_user_06",
			"first_name": "Nikola",
			"last_name": "Jokic",
			"age": 25,
			"created_at": "2021-01-05T10:11:20Z",
			"updated_at": null
		}
	]
	`

	var users []*entities.User

	// json のパース
	if err := json.Unmarshal([]byte(input), &users); err != nil {
		log.Fatal(err)
	}
	log.Printf("users: %v\n", users)

	// ループして、 upsert
	for _, user := range users {
		err := repo.Update(
			context.TODO(),
			user,
		)
		if err != nil {
			log.Fatalln("upsert error:", err)
		}
	}

	// Insert した内容を取得
	res, err := repo.Find(
		context.TODO(),
		bson.M{
			"id": bson.M{
				"$in": bson.A{"test_user_05", "test_user_06"},
			},
		},
	)
	if err != nil {
		log.Fatalln("Find error: ", err)
	}
	log.Printf("Find res: %v\n", res)
	var foundUsers []*entities.User

	data, err := json.Marshal(res)
	if err != nil {
		log.Fatalln("Json Marshal error: ", err)
	}
	if err := json.Unmarshal(data, &foundUsers); err != nil {
		log.Fatalln("Json Unmarshal error: ", err)
	}

	jsonData, err := json.Marshal(foundUsers)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	log.Printf("json: %s", string(jsonData))
	return jsonData, err
}
