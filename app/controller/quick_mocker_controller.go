package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

	var response util.MapString = make(util.MapString)
	var uri string = domain.QuickMockerBooksURL + "/books/" + fmt.Sprint(paramIn.BookID)
	client.GetAPI(uri, func(h *httputil.HTTPResponse) {
		if h.Error != nil {
			response = util.MapString{
				"status":  500,
				"message": h.Message,
				"data":    nil,
			}
			return
		}

		var resData interface{}
		json.Unmarshal(h.Data, &resData)
		response = util.MapString{
			"status":  h.Status,
			"message": h.Message,
			"data":    resData,
		}
	})

	ctx.StatusCode(response["status"].(int))
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
