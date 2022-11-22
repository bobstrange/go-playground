package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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

// Slice を含む構造体のエンコード
func encodeStructWithSlice() {
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

func omitEmpty() {
	// 構造体のタグに omitempty を設定する -> ゼロ値の場合にエンコードされない
	u2 := user2{
		UserID:   "12345",
		UserName: "Jane",
	}
	b, _ := json.Marshal(u2)
	fmt.Println(string(b))
	// 出力の languages は出力されない {"user_id":"12345","user_name":"Jane"}
}

func distinguishNullAndZero() {
	// ゼロ値と区別するためにはポインタを使う (ポインタの場合ゼロ値 nil)
	bottle := bottle{
		Name:  "Coca-Cola",
		Price: 0,
		KCal:  Int(0),
	}
	b, _ := json.Marshal(bottle)
	fmt.Println(string(b))
	// {"name":"Coca-Cola","kcal":0}

}

func disallowUnknownFields() {
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

// JSON の encode
func customMarshal() {
	r := &Record{
		ProcessID: "12345",
		DeletedAt: JSTime(time.Now()),
	}
	b, _ := json.Marshal(r)
	fmt.Println(string(b))
	// unixtime になる {"process_id":"12345","deleted_at":1669128281743}
}

// JSON の decode
func (t *JSTime) UnmarshalJSON(data []byte) error {
	var jsonNumber json.Number
	err := json.Unmarshal(data, &jsonNumber)
	if err != nil {
		return err
	}
	unix, err := jsonNumber.Int64()
	if err != nil {
		return err
	}

	*t = JSTime(time.Unix(0, unix))
	return nil
}

func customUnmarshal() {
	s := `{"process_id":"12345","deleted_at":1669128784280}`
	var r *Record
	if err := json.Unmarshal([]byte(s), &r); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", time.Time(r.DeletedAt).Format(time.RFC3339Nano))
}

func main() {
	encodeStructWithSlice()
	omitEmpty()
	distinguishNullAndZero()
	disallowUnknownFields()
	customMarshal()
	customUnmarshal()
}
