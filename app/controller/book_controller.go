package controller

import (
	"context"
	"net/http"
	"time"

	irisContext "github.com/kataras/iris/v12/context"
	"github.com/rezaif79-ri/iris-api-101/app/domain"
	"github.com/rezaif79-ri/iris-api-101/app/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type bookController struct {
	mongoDb *mongo.Database
}

func NewBookController(db *mongo.Database) domain.BookController {
	return &bookController{
		mongoDb: db,
	}
}

// CreateBook implements domain.BookController.
func (bc *bookController) CreateBook(ctx *irisContext.Context) {
	var b domain.Book
	err := ctx.ReadJSON(&b)
	if err != nil {
		ctx.StopWithJSON(http.StatusConflict, util.RestWrapperObject(http.StatusConflict, "FAIL", util.MapString{
			"error": err.Error(),
		}))
		return
	}

	queryCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := bc.mongoDb.Collection("books").InsertOne(queryCtx,
		bson.D{
			{Key: "author", Value: b.Author},
			{Key: "title", Value: b.Title},
		})
	if err != nil {
		ctx.StopWithJSON(http.StatusConflict, util.RestWrapperObject(http.StatusConflict, "FAIL", util.MapString{
			"error": err.Error(),
		}))
		return
	}

	util.IrisJSONResponse(ctx, http.StatusCreated, "OK", util.MapString{"_id": res.InsertedID})
}

// GetList implements domain.BookController.
func (bc *bookController) GetList(ctx *irisContext.Context) {
	queryCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := bc.mongoDb.Collection("books").Find(queryCtx, bson.D{})
	if err != nil {
		ctx.StopWithJSON(http.StatusConflict, util.RestWrapperObject(http.StatusConflict, "FAIL", util.MapString{
			"error": err.Error(),
		}))
		return
	}
	defer cur.Close(queryCtx)

	var res []util.MapString
	for cur.Next(ctx) {
		var result util.MapString
		err := cur.Decode(&result)
		if err != nil {
			ctx.StopWithJSON(http.StatusConflict, util.RestWrapperObject(http.StatusConflict, "FAIL", util.MapString{
				"error": err.Error(),
			}))
			return
		}
		res = append(res, result)
	}
	if err := cur.Err(); err != nil {
		ctx.StatusCode(http.StatusConflict)
		ctx.JSON(util.RestWrapperObject(http.StatusConflict, "FAIL", err))
		return
	}
	util.IrisJSONResponse(ctx, http.StatusOK, "OK", res)
	// TIP: negotiate the response between server's prioritizes
	// and client's requirements, instead of ctx.JSON:
	// ctx.Negotiation().JSON().MsgPack().Protobuf()
	// ctx.Negotiate(books)
}

// DeleteBook implements domain.BookController.
func (bc *bookController) DeleteBook(ctx *irisContext.Context) {
	queryCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := bc.mongoDb.Collection("books").DeleteOne(queryCtx, bson.D{{Key: "_id", Value: ctx.URLParam("id")}})
	if err != nil {
		ctx.StatusCode(http.StatusConflict)
		ctx.JSON(util.RestWrapperObject(http.StatusConflict, "FAIL", err))
		return
	}

	ctx.StatusCode(http.StatusAccepted)
	ctx.JSON(util.RestWrapperObject(http.StatusAccepted, "OK", util.MapString{
		"count": res.DeletedCount,
	}))
}

// GetOne implements domain.BookController.
func (bc *bookController) GetOne(ctx *irisContext.Context) {
	queryCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var book domain.Book
	err := bc.mongoDb.Collection("books").FindOne(queryCtx, bson.D{
		{Key: "_id", Value: ctx.URLParam("id")},
	}).Decode(&book)
	if err != nil {
		ctx.StatusCode(http.StatusConflict)
		ctx.JSON(util.RestWrapperObject(http.StatusConflict, "FAIL", err))
		return
	}
	util.IrisJSONResponse(ctx, http.StatusOK, "OK", book)
}

// UpdateBook implements domain.BookController.
func (bc *bookController) UpdateBook(ctx *irisContext.Context) {
	var b domain.Book
	err := ctx.ReadJSON(&b)
	if err != nil {
		ctx.StopWithJSON(http.StatusConflict, util.RestWrapperObject(http.StatusConflict, "FAIL", util.MapString{
			"error": err.Error(),
		}))
		return
	}

	queryCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := bc.mongoDb.Collection("books").UpdateByID(queryCtx,
		bson.D{{Key: "_id", Value: ctx.URLParam("id")}},
		bson.D{
			{Key: "author", Value: b.Author},
			{Key: "title", Value: b.Title},
		})
	if err != nil {
		ctx.StopWithJSON(http.StatusConflict, util.RestWrapperObject(http.StatusConflict, "FAIL", util.MapString{
			"error": err.Error(),
		}))
		return
	}

	util.IrisJSONResponse(ctx, http.StatusAccepted, "OK", util.MapString{
		"_id": res.UpsertedID,
	})
}
