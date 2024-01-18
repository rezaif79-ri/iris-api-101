package wkhtmlutil

import (
	"io"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type WkhtmlUtil struct {
}

func (*WkhtmlUtil) NewPDFGenerator() (*wkhtmltopdf.PDFGenerator, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}
	return pdfg, err
}

func (*WkhtmlUtil) NewPDFPageReader(input io.Reader) *wkhtmltopdf.PageReader {
	return wkhtmltopdf.NewPageReader(input)
}

func NewWkhtmlUtil() *WkhtmlUtil {
	return &WkhtmlUtil{}
}
