package main

import (
	"context"
	"os"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
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
	if util.GetEnv("API_DEBUG", "dev") == "dev" {
		app.Logger().SetLevel("debug")
	}
	appLogger := customIrisLog()
	app.UseRouter(appLogger.Handler)
	defer appLogger.Close()

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

func customIrisLog() *accesslog.AccessLog {
	// Initialize a new access log middleware.
	ac := accesslog.New(os.Stdout)
	// The default configuration:
	ac.Delim = '|'
	ac.TimeFormat = "2006-01-02 15:04:05"
	ac.Async = false
	ac.IP = true
	ac.BytesReceivedBody = true
	ac.BytesSentBody = true
	ac.BytesReceived = false
	ac.BytesSent = false
	ac.BodyMinify = true
	ac.RequestBody = true
	ac.ResponseBody = false
	ac.KeepMultiLineError = true
	ac.PanicLog = accesslog.LogHandler

	// Default line format if formatter is missing:
	// Time|Latency|Code|Method|Path|IP|Path Params Query Fields|Bytes Received|Bytes Sent|Request|Response|
	//
	// Set Custom Formatter:
	ac.SetFormatter(&accesslog.JSON{
		Indent:    "  ",
		HumanTime: true,
	})
	// ac.SetFormatter(&accesslog.CSV{})
	// ac.SetFormatter(&accesslog.Template{Text: "{{.Code}}"})
	return ac
}
