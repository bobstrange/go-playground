package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
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

	// 想定していないフィールドが JSON に含まれていた時にエラーを出す
	// DisallowUnknownFields

	type Rectangle struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	}

	s := []byte(`{"width": 10, "height": 20, "radius": 5}`)
	var rect Rectangle
	d := json.NewDecoder(bytes.NewReader(s))
	d.DisallowUnknownFields()
	if err := d.Decode(&rect); err != nil {
		fmt.Println(err) // Will print `json: unknown field "radius"`
	}

	customMarshal()

}

// Marshal と Unmarshal を拡張する
type Record struct {
	ProcessID string `json:"process_id"`
	DeletedAt JSTime `json:"deleted_at"`
}

type JSTime time.Time

func (t JSTime) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)
	if tt.IsZero() {
		return []byte("null"), nil
	}
	v := strconv.Itoa(int(tt.UnixMilli()))
	return []byte(v), nil
}

func customMarshal() {
	r := &Record{
		ProcessID: "12345",
		DeletedAt: JSTime(time.Now()),
	}
	b, _ := json.Marshal(r)
	fmt.Println(string(b))

}
