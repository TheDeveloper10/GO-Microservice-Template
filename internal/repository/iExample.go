package repository

import "msvc-template/internal/util"

type IExample interface {
	Func() util.StatusCode
}

func NewExampleRepository() IExample {
	return &Example{}
}
