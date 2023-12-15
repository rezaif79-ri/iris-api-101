package router

import (
	"github.com/kataras/iris/v12"
	"github.com/rezaif79-ri/iris-api-101/app/controller"
)

func SetupIrisRouter(app *iris.Application) *iris.Application {
	// Init controller
	bookController := controller.NewBookController()

	books := app.Party("books")
	books.Use(iris.Compression)
	books.Get("", bookController.GetList)
	books.Post("", bookController.CreateBook)

	return app
}
