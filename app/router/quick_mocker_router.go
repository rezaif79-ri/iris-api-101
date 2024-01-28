package router

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/rezaif79-ri/iris-api-101/app/controller"
)

func AddQuickMockerBookRouter(group router.Party) {
	// Init file zipper controller
	mocker := controller.NewQuickMockerController()

	group.Use(iris.Compression)
	group.Get("", mocker.GetListBooks)
	group.Get("{id}", mocker.GetBookDetail)
	group.Post("", mocker.InsertBook)
	group.Put("{id}", mocker.UpdateBook)
	group.Delete("{id}", mocker.DeleteBook)
}
