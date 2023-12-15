package main

import (
	"github.com/kataras/iris/v12"
	"github.com/rezaif79-ri/iris-api-101/app/router"
)

func main() {
	app := iris.New()

	router.SetupIrisRouter(app)

	app.Listen("127.0.0.1:8080")
}
