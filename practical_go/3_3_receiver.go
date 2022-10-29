package main

import "fmt"

type Person struct {
	FirstName string
}

// レシーバーを値型にすると、レシーバーの属性の変更が反映されない (値渡しなので)
func (p Person) NotSetFirstName(name string) {
	p.FirstName = name
}

// ポインタ型にすると、レシーバーの属性の変更が反映される
func (p *Person) SetFirstName(name string) {
	p.FirstName = name
}

//
func main() {
	user := &Person{
		FirstName: "John",
	}
	user.NotSetFirstName("Jane")
	fmt.Println("user: ", user)

	user.SetFirstName("Smith")
	fmt.Println("user: ", user)
}
