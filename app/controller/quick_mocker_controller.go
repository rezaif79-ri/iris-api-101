package controller

import (
	"github.com/kataras/iris/v12/context"
	"github.com/rezaif79-ri/iris-api-101/app/domain"
)

type QuickMockerControllerImpl struct {
}

// DeleteBook implements domain.QuickMockerController.
func (*QuickMockerControllerImpl) DeleteBook(*context.Context) {
	panic("unimplemented")
}

// GetBookDetail implements domain.QuickMockerController.
func (*QuickMockerControllerImpl) GetBookDetail(*context.Context) {
	panic("unimplemented")
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
