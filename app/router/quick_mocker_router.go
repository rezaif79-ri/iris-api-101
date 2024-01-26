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
	group.Post("", mocker.GetListBooks)
	group.Post("{id}", mocker.GetBookDetail)
	group.Put("{id}", mocker.UpdateBook)
	group.Delete("{id}", mocker.DeleteBook)
}
