package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/rezaif79-ri/iris-api-101/app/domain"
	"github.com/rezaif79-ri/iris-api-101/app/util"
	httputil "github.com/rezaif79-ri/iris-api-101/app/util/http_util"
)

type QuickMockerControllerImpl struct {
}

// DeleteBook implements domain.QuickMockerController.
func (*QuickMockerControllerImpl) DeleteBook(*context.Context) {
	panic("unimplemented")
}

// GetBookDetail implements domain.QuickMockerController.
func (qmc *QuickMockerControllerImpl) GetBookDetail(ctx *context.Context) {
	type BookDetailParams struct {
		BookID int `param:"id"`
	}
	var paramIn BookDetailParams
	if err := ctx.ReadParams(&paramIn); err != nil {
		ctx.StopWithJSON(http.StatusConflict, util.RestWrapperObject(http.StatusConflict, "FAIL", util.MapString{
			"error": err.Error(),
		}))
		return
	}
	client := httputil.NewHTTPClient(time.Second * 10)

	response := httputil.HTTPResponse{}

	var uri string = domain.QuickMockerBooksURL + "/books/" + fmt.Sprint(paramIn.BookID)
	client.GetAPI(uri, func(r *http.Response, err error) {
		if err != nil {
			ctx.StatusCode(r.StatusCode)
			ctx.JSON(iris.Map{
				"error": err.Error(),
			})
			return
		}

		var res map[string]interface{}
		body, err := httputil.ConvertBodyToBytes(r.Body)
		if err != nil {
			ctx.StatusCode(500)
			ctx.JSON(iris.Map{
				"error": err.Error(),
			})
			return
		}
		json.Unmarshal(body, &res)
		response = httputil.HTTPResponse{
			Status:  r.StatusCode,
			Message: "OK",
			Error:   nil,
			Data:    res,
		}
	})

	ctx.StatusCode(200)
	ctx.JSON(response)
}

// GetListBooks implements domain.QuickMockerController.
func (*QuickMockerControllerImpl) GetListBooks(*context.Context) {
	panic("unimplemented")
}

// InsertBook implements domain.QuickMockerController.
func (*QuickMockerControllerImpl) InsertBook(*context.Context) {
	panic("unimplemented")
}

// UpdateBook implements domain.QuickMockerController.
func (*QuickMockerControllerImpl) UpdateBook(*context.Context) {
	panic("unimplemented")
}

func NewQuickMockerController() domain.QuickMockerController {
	return &QuickMockerControllerImpl{}
}
