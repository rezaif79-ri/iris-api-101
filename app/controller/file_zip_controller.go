package controller

import (
	"archive/zip"
	"bytes"
	"io"
	"mime"

	"github.com/kataras/iris/v12/context"
	"github.com/rezaif79-ri/iris-api-101/app/domain"
	"github.com/rezaif79-ri/iris-api-101/app/util"
)

type FileZipControllerImpl struct {
}

func NewFileZipController() domain.FileZipController {
	return &FileZipControllerImpl{}
}

// ZipOneFile implements domain.FileZipController.
func (*FileZipControllerImpl) ZipOneFile(ctx *context.Context) {
	fileIn, headerIn, err := ctx.FormFile("file")
	if err != nil {
		ctx.StopWithJSON(409, util.RestWrapperObject(409, "FAIL", util.MapString{
			"status":  409,
			"message": err.Error(),
			"data":    err,
		}))
		return
	}

	writerFileBuf := bytes.NewBuffer(nil)
	var zipWriter = zip.NewWriter(writerFileBuf)

	fileType, err := mime.ExtensionsByType(headerIn.Header.Get("Content-Type"))
	if err != nil {
		ctx.StopWithJSON(409, util.RestWrapperObject(409, "FAIL", util.MapString{
			"status":  409,
			"message": err.Error(),
			"data":    err,
		}))
		return
	}

	fileName := headerIn.Filename + fileType[0]
	fileWriter, err := zipWriter.Create(fileName)
	if err != nil {
		ctx.StopWithJSON(500, util.MapString{
			"status":  500,
			"message": err.Error(),
			"data":    err,
		})
		return
	}
	if _, err := io.Copy(fileWriter, fileIn); err != nil {
		ctx.StopWithJSON(500, util.MapString{
			"status":  500,
			"message": err.Error(),
			"data":    err,
		})
		return
	}
	if err := zipWriter.Close(); err != nil {
		ctx.StopWithJSON(500, util.MapString{
			"status":  500,
			"message": err.Error(),
			"data":    err,
		})
		return
	}

	var dummyFileName = "dummy.zip"

	ctx.Header("Content-type", "application/zip")
	ctx.Header("Content-Disposition", "attachment; filename=\""+dummyFileName+"\"")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Write(writerFileBuf.Bytes())
	writerFileBuf.Reset()
}
