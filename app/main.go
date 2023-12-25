package main

import (
	"context"

	"github.com/kataras/iris/v12"
	"github.com/rezaif79-ri/iris-api-101/app/config"
	"github.com/rezaif79-ri/iris-api-101/app/router"
)

func main() {
	// Init mongo db
	mongoDb, err := config.OpenMongoBookDB()
	if err != nil {
		panic("setup mongo conn error: " + err.Error())
	}

	app := iris.New()

	router.SetupIrisRouter(app, mongoDb)

	app.Listen("127.0.0.1:8080")
	defer func() {
		if err = mongoDb.Client().Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
}
