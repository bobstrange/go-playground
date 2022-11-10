package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ip struct {
	Origin string `json:"origin"`
	URL    string `json:"url"`
}

func main() {
	f, err := os.Open("ip.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var resp ip
	// 入力が io.Reader interface を満たす場合 os.Stdin や http.Response.Body など
	if err := json.NewDecoder(f).Decode(&resp); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
	// {Origin:255.255.255.255 URL:https://httpbin.org/get}

	s := `{"origin": "255.255.255.255", "url": "https://httpbin.org/get"}`
	var resp2 ip
	if err = json.Unmarshal([]byte(s), &resp2); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", resp2)

	var b bytes.Buffer
	ip := ip{
		Origin: "255.255.255.255",
		URL:    "https://httpbin.org/get",
	}
	_ = json.NewEncoder(&b).Encode(ip)
	fmt.Printf("%v\n", b.String())
	b2, _ := json.Marshal(ip)
	fmt.Println(string(b2))
}
