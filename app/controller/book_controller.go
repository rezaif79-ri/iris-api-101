package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/rezaif79-ri/iris-api-101/app/domain"
)

type bookController struct {
}

func NewBookController() domain.BookController {
	return &bookController{}
}

// CreateBook implements domain.BookController.
func (bookController) CreateBook(ctx *context.Context) {
	var b domain.Book
	err := ctx.ReadJSON(&b)
	// TIP: use ctx.ReadBody(&b) to bind
	// any type of incoming data instead.
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Book creation failure").DetailErr(err))
		// TIP: use ctx.StopWithError(code, err) when only
		// plain text responses are expected on errors.
		return
	}

	println("Received Book: " + b.Title)

	ctx.StatusCode(iris.StatusCreated)
}

// GetList implements domain.BookController.
func (bookController) GetList(ctx *context.Context) {
	books := []domain.Book{
		{
			Title:  "Mastering Concurrency in Go",
			Author: "John Doe",
		},
		{
			Title:  "Go Design Patterns",
			Author: "John Doe",
		},
		{
			Title:  "Black Hat Go",
			Author: "John Doe",
		},
	}

	ctx.JSON(books)
	// TIP: negotiate the response between server's prioritizes
	// and client's requirements, instead of ctx.JSON:
	// ctx.Negotiation().JSON().MsgPack().Protobuf()
	// ctx.Negotiate(books)
}
