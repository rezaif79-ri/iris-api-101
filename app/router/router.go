package router

import (
	"context"

	"github.com/kataras/iris/v12"
	"github.com/rezaif79-ri/iris-api-101/app/config"
	"github.com/rezaif79-ri/iris-api-101/app/controller"
)

func SetupIrisRouter(app *iris.Application) *iris.Application {
	// Init mongo db
	mongoDb, err := config.SetupMongoConn()
	if err != nil {
		panic("setup mongo conn error: " + err.Error())
	}
	defer func() {
		if err = mongoDb.Client().Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	// Init controller
	bookController := controller.NewBookController(mongoDb)

	books := app.Party("books")
	books.Use(iris.Compression)
	books.Get("", bookController.GetList)
	books.Post("", bookController.CreateBook)

	return app
}
