package router

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/rezaif79-ri/iris-api-101/app/controller"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddBookRouter(group router.Party, mongoDb *mongo.Database) {
	// Init controller
	bookController := controller.NewBookController(mongoDb)
	group.Use(iris.Compression)
	group.Get("", bookController.GetList)
	group.Post("", bookController.CreateBook)
	group.Put("{id}", bookController.UpdateBook)
	group.Get("{id}", bookController.GetOne)
	group.Delete("{id}", bookController.DeleteBook)
}
