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
		util.IrisJSONResponse(ctx, http.StatusConflict, "FAIL", util.MapString{
			"error": err.Error(),
		})
		return
	}
	client := httputil.NewHTTPClient(time.Second * 10)

	var uri string = domain.QuickMockerBooksURL + "/books/" + fmt.Sprint(paramIn.BookID)
	client.GetAPI(uri, func(h *httputil.HTTPResponse) {
		if h.Error != nil {
			util.IrisJSONResponse(ctx, 500, h.Message, nil)
			return
		}
		var resData interface{}
		json.Unmarshal(h.Data, &resData)
		util.IrisJSONResponse(ctx, h.Status, h.Message, resData)
		return
	})
	return
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
