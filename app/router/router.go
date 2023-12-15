package router

import "github.com/kataras/iris/v12"

func SetupIrisRouter(app *iris.Application) *iris.Application {
	books := app.Party("books")
	books.Use(iris.Compression)
	return app
}
