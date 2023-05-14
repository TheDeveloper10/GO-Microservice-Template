package service

import (
	"msvc-template/internal/repository"
	"msvc-template/internal/util"
)

type Example struct {
	repo repository.IExample
}

func (svc *Example) Func() util.StatusCode {
	return svc.repo.Func()
}
