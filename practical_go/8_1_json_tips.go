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

type user2 struct {
	UserID    string   `json:"user_id"`
	UserName  string   `json:"user_name"`
	Languages []string `json:"languages,omitempty"`
}

type bottle struct {
	Name  string `json:"name"`
	Price int    `json:"price,omitempty"`
	KCal  *int   `json:"kcal,omitempty"`
}

func Int(v int) *int { return &v }

func main() {
	// Slice を含む構造体のエンコード
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

	// 構造体のタグに omitempty を設定する -> ゼロ値の場合にエンコードされない
	u2 := user2{
		UserID:   "12345",
		UserName: "Jane",
	}
	b, _ = json.Marshal(u2)
	fmt.Println(string(b))
	// 出力の languages は出力されない {"user_id":"12345","user_name":"Jane"}

	// ゼロ値と区別するためにはポインタを使う (ポインタの場合ゼロ値 nil)
	bottle := bottle{
		Name:  "Coca-Cola",
		Price: 0,
		KCal:  Int(0),
	}
	b, _ = json.Marshal(bottle)
	fmt.Println(string(b))
	// {"name":"Coca-Cola","kcal":0}
}
