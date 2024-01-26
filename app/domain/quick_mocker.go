package domain

import "github.com/kataras/iris/v12/context"

const QuickMockerBooksURL string = "https://9k8lv1chxz.api.quickmocker.com"

type QuickMockerController interface {
	GetBookDetail(*context.Context)
	GetListBooks(*context.Context)
	InsertBook(*context.Context)
	UpdateBook(*context.Context)
	DeleteBook(*context.Context)
}
