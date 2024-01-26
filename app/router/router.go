package router

import (
	"bytes"
	"fmt"
	"io"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/rezaif79-ri/iris-api-101/app/util"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupIrisRouter(app *iris.Application, mongoDb *mongo.Database) *iris.Application {
	books := app.Party("books/")
	files := app.Party("files/")
	bookMocker := app.Party("mocker/books/")

	AddBookRouter(books, mongoDb)
	AddFileZipperRouter(files)
	AddQuickMockerBookRouter(bookMocker)

	files.Post("", func(ctx *context.Context) {
		inputFile, inputHeader, err := ctx.FormFile("input_file")
		if err != nil {
			ctx.StopWithJSON(409, util.RestWrapperObject(409, "FAIL", util.MapString{
				"error1": err.Error(),
			}))
			return
		}

		fmt.Println("files input:", inputFile)
		fmt.Println("files header content type :", inputHeader.Header.Get("Content-Type"))
		fileBuf := bytes.NewBuffer(nil)
		if _, err := io.Copy(fileBuf, inputFile); err != nil {
			ctx.StopWithJSON(409, util.RestWrapperObject(409, "FAIL", util.MapString{
				"error2": err.Error(),
			}))
			return
		}
		fmt.Println("bytes: ", fileBuf.Bytes())

		ctx.StatusCode(200)
		ctx.JSON(util.MapString{
			"message": "OK",
		})
	})

	return app
}
