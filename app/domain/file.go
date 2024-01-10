package domain

import (
	"mime/multipart"

	"github.com/kataras/iris/v12"
)

type File struct {
	InputFile multipart.FileHeader `form:"input_file"`
	FileName  string               `form:"file_name"`
}

type FileZipController interface {
	ZipOneFile(iris.Context)
	ZipMultiFile(iris.Context)
}
