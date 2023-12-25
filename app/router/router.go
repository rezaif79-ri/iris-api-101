package router

import (
	"github.com/kataras/iris/v12"
	"github.com/rezaif79-ri/iris-api-101/app/controller"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupIrisRouter(app *iris.Application, mongoDb *mongo.Database) *iris.Application {
	// Init controller
	bookController := controller.NewBookController(mongoDb)

	books := app.Party("books")
	books.Use(iris.Compression)
	books.Get("", bookController.GetList)
	books.Post("", bookController.CreateBook)

	return app
}
