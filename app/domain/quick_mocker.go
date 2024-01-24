package domain

import "github.com/kataras/iris/v12/context"

type QuickMocker interface {
	GetBookDetail(*context.Context)
	GetListBooks(*context.Context)
	InsertBook(*context.Context)
	UpdateBook(*context.Context)
	DeleteBook(*context.Context)
}
