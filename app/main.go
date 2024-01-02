package main

import (
	"context"

	"github.com/kataras/iris/v12"
	"github.com/rezaif79-ri/iris-api-101/app/config"
	"github.com/rezaif79-ri/iris-api-101/app/router"
	"github.com/rezaif79-ri/iris-api-101/app/util"
)

func main() {
	// Init mongo db
	mongoDb, err := config.OpenMongoBookDB()
	if err != nil {
		panic("setup mongo conn error: " + err.Error())
	}

	app := iris.New()

	router.SetupIrisRouter(app, mongoDb)

	appUrl := util.GetEnv("API_URL", "localhost")
	appPort := util.GetEnv("API_PORT", "8080")
	app.Listen(appUrl + ":" + appPort)
	defer func() {
		if err = mongoDb.Client().Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
}
