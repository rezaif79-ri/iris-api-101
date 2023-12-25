package controller

import (
	"context"
	"net/http"

	"github.com/kataras/iris/v12"
	irisContext "github.com/kataras/iris/v12/context"
	"github.com/rezaif79-ri/iris-api-101/app/domain"
	"github.com/rezaif79-ri/iris-api-101/app/util"
	"go.mongodb.org/mongo-driver/mongo"
)

type bookController struct {
	mongoDb *mongo.Database
}

// DeleteBook implements domain.BookController.
func (*bookController) DeleteBook(ctx *irisContext.Context) {
	panic("unimplemented")
}

// GetOne implements domain.BookController.
func (*bookController) GetOne(ctx *irisContext.Context) {
	panic("unimplemented")
}

// UpdateBook implements domain.BookController.
func (*bookController) UpdateBook(ctx *irisContext.Context) {
	panic("unimplemented")
}

func NewBookController(db *mongo.Database) domain.BookController {
	return &bookController{
		mongoDb: db,
	}
}

// CreateBook implements domain.BookController.
func (bookController) CreateBook(ctx *irisContext.Context) {
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
func (bc *bookController) GetList(ctx *irisContext.Context) {
	books := bc.mongoDb.Collection("books")

	res, err := books.Find(context.Background(), nil)
	if err != nil {
		ctx.StopWithJSON(http.StatusConflict, nil)
	}

	ctx.StatusCode(http.StatusOK)
	ctx.JSON(util.RestWrapperObject(http.StatusOK, "OK", res))
	// TIP: negotiate the response between server's prioritizes
	// and client's requirements, instead of ctx.JSON:
	// ctx.Negotiation().JSON().MsgPack().Protobuf()
	// ctx.Negotiate(books)
}
