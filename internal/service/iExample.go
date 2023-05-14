package service

import (
	"msvc-template/internal/repository"
	"msvc-template/internal/util"
)

type IExample interface {
	Func() util.StatusCode
}

func NewExampleService(repo repository.IExample) IExample {
	return &Example{
		repo: repo,
	}
}
