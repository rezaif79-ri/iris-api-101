package controller

import (
	"bytes"
	"net/http"
	"time"

	"github.com/kataras/iris/v12/context"
	"github.com/rezaif79-ri/iris-api-101/app/domain"
	"github.com/rezaif79-ri/iris-api-101/app/util"
	wkhtmlutil "github.com/rezaif79-ri/iris-api-101/app/util/wkhtml_util"
)

type WkhtmlControllerImpl struct {
	*wkhtmlutil.WkhtmlUtil
}

func NewWkhtmlController(wkhtmlutil *wkhtmlutil.WkhtmlUtil) domain.WkhtmlController {
	return &WkhtmlControllerImpl{
		WkhtmlUtil: wkhtmlutil,
	}
}

// ConvertHtmlToPDF implements domain.WkhtmlController.
func (wci *WkhtmlControllerImpl) ConvertHtmlToPDF(ctx *context.Context) {
	var b domain.HtmlConvertPDF
	err := ctx.ReadJSON(&b)
	if err != nil {
		ctx.StopWithJSON(http.StatusConflict, util.RestWrapperObject(http.StatusConflict, "FAIL", util.MapString{
			"error": err.Error(),
		}))
		return
	}

	pdfg, err := wci.NewPDFGenerator()
	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, util.RestWrapperObject(
			http.StatusInternalServerError, "FAIL",
			util.MapString{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
				"data":    err,
			}))
		return
	}
	dataBuffer := bytes.NewBuffer(b.Data)
	// Add to document
	pdfg.AddPage(wci.NewPDFPageReader(dataBuffer))

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, util.RestWrapperObject(
			http.StatusInternalServerError, "FAIL",
			util.MapString{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
				"data":    err,
			}))
		return
	}

	pdfData := bytes.NewReader(pdfg.Bytes())
	ctx.ServeContent(pdfData, "test", time.Now())
}

// ConvertMultiHtmlToPDF implements domain.WkhtmlController.
func (wci *WkhtmlControllerImpl) ConvertMultiHtmlToPDF(ctx *context.Context) {
	var b domain.MultiHtmlConvertPDF
	err := ctx.ReadJSON(&b)
	if err != nil {
		ctx.StopWithJSON(http.StatusConflict, util.RestWrapperObject(http.StatusConflict, "FAIL", util.MapString{
			"error": err.Error(),
		}))
		return
	}

	pdfg, err := wci.NewPDFGenerator()
	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, util.RestWrapperObject(
			http.StatusInternalServerError, "FAIL",
			util.MapString{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
				"data":    err,
			}))
		return
	}

	for i := range b.Data {
		// Add to document
		pdfg.AddPage(wci.NewPDFPageReader(bytes.NewBuffer(b.Data[i])))

	}

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, util.RestWrapperObject(
			http.StatusInternalServerError, "FAIL",
			util.MapString{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
				"data":    err,
			}))
		return
	}

	pdfData := bytes.NewReader(pdfg.Bytes())
	ctx.ServeContent(pdfData, "test", time.Now())
}
