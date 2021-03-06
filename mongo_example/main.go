package main

import (
	"github.com/bobstrange/go-playground/mongo_example/handler"
	"github.com/bobstrange/go-playground/mongo_example/repositories"
)

func main() {
	repo := repositories.NewRepo()
	handler.Handle(repo)
}
