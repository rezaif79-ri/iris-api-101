package router

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/rezaif79-ri/iris-api-101/app/controller"
)

func AddFileZipperRouter(group router.Party) {
	// Init file zipper controller
	zipper := controller.NewFileZipController()

	group.Use(iris.Compression)
	group.Post("zipper/single", zipper.ZipOneFile)
	group.Post("zipper/multi", zipper.ZipMultiFile)
}
