package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	UserID    string   `json:"user_id"`
	UserName  string   `json:"user_name"`
	Languages []string `json:"languages"`
}

func main() {

	// Languages に空スライスを設定しない場合
	u := user{
		UserID:   "12345",
		UserName: "John",
	}

	b, _ := json.Marshal(u)
	fmt.Println(string(b))
	// 出力の languages は null になる {"user_id":"12345","user_name":"John","languages":null}

	// Languages に空スライスを設定すると
	u = user{
		UserID:    "12345",
		UserName:  "John",
		Languages: []string{},
	}
	b, _ = json.Marshal(u)
	fmt.Println(string(b))
	// 出力の languages は [] になる {"user_id":"12345","user_name":"John","languages":[]}

}
