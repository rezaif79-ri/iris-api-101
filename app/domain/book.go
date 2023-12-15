package domain

import "github.com/kataras/iris/v12"

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookController interface {
	GetList(ctx iris.Context)
	CreateBook(ctx iris.Context)
}
