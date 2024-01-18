package controller

import (
	"github.com/kataras/iris/v12/context"
	"github.com/rezaif79-ri/iris-api-101/app/domain"
)

type WkhtmlControllerImpl struct {
}

func NewWkhtmlController() domain.WkhtmlController {
	return &WkhtmlControllerImpl{}
}

// ConvertHtmlToPDF implements domain.WkhtmlController.
func (*WkhtmlControllerImpl) ConvertHtmlToPDF(*context.Context) {
	panic("unimplemented")
}
