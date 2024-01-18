package wkhtmlutil

import (
	"io"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type wkhtmlUtil struct {
}

func (*wkhtmlUtil) NewPDFGenerator() (*wkhtmltopdf.PDFGenerator, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}
	return pdfg, err
}

func (*wkhtmlUtil) NewPDFPageReader(input io.Reader) *wkhtmltopdf.PageReader {
	return wkhtmltopdf.NewPageReader(input)
}

func NewWkhtmlUtil() *wkhtmlUtil {
	return &wkhtmlUtil{}
}
