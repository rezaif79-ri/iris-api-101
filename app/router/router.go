package router

import (
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupIrisRouter(app *iris.Application, mongoDb *mongo.Database) *iris.Application {
	books := app.Party("books")
	AddBookRouter(books, mongoDb)

	return app
}
