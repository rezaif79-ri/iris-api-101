package controller

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/rezaif79-ri/iris-api-101/app/domain"
)

type bookController struct {
}

// DeleteBook implements domain.BookController.
func (*bookController) DeleteBook(ctx *context.Context) {
	panic("unimplemented")
}

// GetOne implements domain.BookController.
func (*bookController) GetOne(ctx *context.Context) {
	panic("unimplemented")
}

// UpdateBook implements domain.BookController.
func (*bookController) UpdateBook(ctx *context.Context) {
	panic("unimplemented")
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

	ctx.StatusCode(http.StatusOK)
	ctx.JSON(map[string]interface{}{
		"status":  http.StatusOK,
		"message": "OK",
		"data":    books,
	})
	// TIP: negotiate the response between server's prioritizes
	// and client's requirements, instead of ctx.JSON:
	// ctx.Negotiation().JSON().MsgPack().Protobuf()
	// ctx.Negotiate(books)
}
