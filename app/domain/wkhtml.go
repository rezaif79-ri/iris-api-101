package domain

import "github.com/kataras/iris/v12"

type HtmlConvertPDF struct {
	Data []byte `json:"data"`
}

type WkhtmlController interface {
	ConvertHtmlToPDF(iris.Context)
}
