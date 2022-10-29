package main

import "fmt"

type Book struct {
	Title string
	ISBN  string
}

func (b Book) GetAmazonURL() string {
	return "https://amazon.co.jp/dp/" + b.ISBN
}

type OreillyBook struct {
	// 構造体の埋め込み
	// Book Book 型名 Book と同じ名前のフィールドがあるのとほぼ同じ
	Book
	ISBN13 string
}

func (o OreillyBook) GetOreillyURL() string {
	return "https://www.oreilly.co.jp/books/" + o.ISBN13 + "/"
}

func main() {
	book := OreillyBook{
		ISBN13: "9784873119038",
		Book: Book{
			Title: "Real World HTTP",
			ISBN:  "4873119030",
		},
	}

	fmt.Println(book.GetAmazonURL())
	fmt.Println(book.GetOreillyURL())
	fmt.Println(book.Book.GetAmazonURL())
}
